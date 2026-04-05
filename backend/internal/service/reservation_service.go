package service

import (
	"context"
	"errors"
	"time"

	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
)

var (
	ErrReservationNotFound = errors.New("rezervasyon bulunamadı")
	ErrBookNotAvailable    = errors.New("kitap müsait değil")
	ErrNotOwnerOrAdmin     = errors.New("bu rezervasyona erişim yetkiniz bulunmuyor")
	ErrAlreadyReturned     = errors.New("bu kitap zaten iade edilmiş")
	ErrInvalidDueDate      = errors.New("geçersiz iade tarihi")
	ErrMaxReservations     = errors.New("maksimum aktif rezervasyon sayısına ulaşıldı")
)

const MaxActiveReservationsPerUser = 5

type ReservationService struct {
	reservationRepo *repository.ReservationRepository
	bookRepo        *repository.BookRepository
	userRepo        *repository.UserRepository
}

func NewReservationService(
	reservationRepo *repository.ReservationRepository,
	bookRepo *repository.BookRepository,
	userRepo *repository.UserRepository,
) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationRepo,
		bookRepo:        bookRepo,
		userRepo:        userRepo,
	}
}

func (s *ReservationService) Create(ctx context.Context, userID string, req dto.CreateReservationRequest) (*dto.ReservationResponse, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return nil, ErrInvalidDueDate
	}
	if dueDate.Before(time.Now()) {
		return nil, ErrInvalidDueDate
	}

	book, err := s.bookRepo.GetByID(ctx, req.BookID)
	if err != nil {
		return nil, ErrBookNotFound
	}
	if !book.Available() {
		return nil, ErrBookNotAvailable
	}

	activeCount, err := s.reservationRepo.CountActiveByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if activeCount >= MaxActiveReservationsPerUser {
		return nil, ErrMaxReservations
	}

	reservation := &model.Reservation{
		UserID:  userID,
		BookID:  req.BookID,
		Status:  model.StatusActive,
		DueDate: dueDate,
	}

	if err := s.reservationRepo.Create(ctx, reservation); err != nil {
		return nil, err
	}

	book.AvailableCopies--
	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, err
	}

	return toReservationResponse(reservation), nil
}

func (s *ReservationService) GetByID(ctx context.Context, id, userID, userRole string) (*dto.ReservationResponse, error) {
	reservation, err := s.reservationRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrReservationNotFound
	}

	if userRole != string(model.RoleAdmin) && reservation.UserID != userID {
		return nil, ErrNotOwnerOrAdmin
	}

	return toReservationResponse(reservation), nil
}

func (s *ReservationService) Return(ctx context.Context, id, userID, userRole string) error {
	reservation, err := s.reservationRepo.GetByID(ctx, id)
	if err != nil {
		return ErrReservationNotFound
	}

	if userRole != string(model.RoleAdmin) && reservation.UserID != userID {
		return ErrNotOwnerOrAdmin
	}

	if reservation.Status == model.StatusReturned {
		return ErrAlreadyReturned
	}

	now := time.Now()
	reservation.Status = model.StatusReturned
	reservation.ReturnedAt = &now

	if err := s.reservationRepo.Update(ctx, reservation); err != nil {
		return err
	}

	book, err := s.bookRepo.GetByID(ctx, reservation.BookID)
	if err != nil {
		return err
	}
	book.AvailableCopies++
	return s.bookRepo.Update(ctx, book)
}

func (s *ReservationService) ListByUserID(ctx context.Context, targetUserID, requestingUserID, userRole string, q dto.UserReservationsQuery) ([]dto.ReservationResponse, int, error) {
	if userRole != string(model.RoleAdmin) && targetUserID != requestingUserID {
		return nil, 0, ErrNotOwnerOrAdmin
	}

	_, err := s.userRepo.GetByID(ctx, targetUserID)
	if err != nil {
		return nil, 0, errors.New("kullanıcı bulunamadı")
	}

	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 || q.Limit > 50 {
		q.Limit = 10
	}

	reservations, totalCount, err := s.reservationRepo.ListByUserID(ctx, targetUserID, q.Status, q.Page, q.Limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.ReservationResponse
	for _, r := range reservations {
		responses = append(responses, *toReservationResponse(&r))
	}

	return responses, totalCount, nil
}

func toReservationResponse(r *model.Reservation) *dto.ReservationResponse {
	return &dto.ReservationResponse{
		ID:         r.ID,
		UserID:     r.UserID,
		BookID:     r.BookID,
		Status:     string(r.Status),
		ReservedAt: r.ReservedAt,
		DueDate:    r.DueDate,
		ReturnedAt: r.ReturnedAt,
	}
}
