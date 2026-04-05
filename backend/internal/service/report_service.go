package service

import (
	"context"
	"time"

	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
)

type ReportService struct {
	reservationRepo *repository.ReservationRepository
	bookRepo        *repository.BookRepository
	userRepo        *repository.UserRepository
}

func NewReportService(
	reservationRepo *repository.ReservationRepository,
	bookRepo *repository.BookRepository,
	userRepo *repository.UserRepository,
) *ReportService {
	return &ReportService{
		reservationRepo: reservationRepo,
		bookRepo:        bookRepo,
		userRepo:        userRepo,
	}
}

func (s *ReportService) Generate(ctx context.Context, q dto.ReportQuery) (*dto.ReportResponse, error) {
	reportType := q.Type
	if reportType == "" {
		reportType = "reservations"
	}

	var from, to time.Time
	var err error

	if q.From != "" {
		from, err = time.Parse("2006-01-02", q.From)
		if err != nil {
			from = time.Now().AddDate(0, -1, 0)
		}
	} else {
		from = time.Now().AddDate(0, -1, 0)
	}

	if q.To != "" {
		to, err = time.Parse("2006-01-02", q.To)
		if err != nil {
			to = time.Now()
		}
	} else {
		to = time.Now()
	}

	response := &dto.ReportResponse{
		GeneratedAt: time.Now(),
		Type:        reportType,
	}

	switch reportType {
	case "reservations":
		total, _ := s.reservationRepo.CountByDateRange(ctx, from, to)
		active, _ := s.reservationRepo.CountActiveByDateRange(ctx, from, to)
		returned, _ := s.reservationRepo.CountReturnedByDateRange(ctx, from, to)

		response.Summary = dto.ReportSummary{
			Total:    total,
			Active:   active,
			Returned: returned,
		}

		topBooks, _ := s.reservationRepo.TopReservedBooks(ctx, 10)
		for _, tb := range topBooks {
			response.TopBooks = append(response.TopBooks, dto.TopBookReport{
				BookID:           tb.BookID,
				Title:            tb.Title,
				ReservationCount: tb.ReservationCount,
			})
		}

		details, _ := s.reservationRepo.RecentReservationsWithUsers(ctx, 50)
		for _, d := range details {
			response.Reservations = append(response.Reservations, dto.ReservationDetailReport{
				BookTitle:  d.BookTitle,
				UserName:   d.UserName,
				UserEmail:  d.UserEmail,
				Status:     d.Status,
				ReservedAt: d.ReservedAt,
				DueDate:    d.DueDate,
			})
		}

	case "books":
		total, _ := s.bookRepo.Count(ctx)
		response.Summary = dto.ReportSummary{Total: total}

	case "users":
		total, _ := s.userRepo.Count(ctx)
		activeRes, _ := s.reservationRepo.CountByStatus(ctx, model.StatusActive)
		returnedRes, _ := s.reservationRepo.CountByStatus(ctx, model.StatusReturned)
		response.Summary = dto.ReportSummary{
			Total:    total,
			Active:   activeRes,
			Returned: returnedRes,
		}
	}

	return response, nil
}
