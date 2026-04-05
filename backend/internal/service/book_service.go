package service

import (
	"context"
	"errors"

	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
)

var (
	ErrBookNotFound        = errors.New("kitap bulunamadı")
	ErrISBNAlreadyExists   = errors.New("bu ISBN numarası zaten kayıtlı")
	ErrBookHasReservations = errors.New("aktif rezervasyonu olan kitap silinemez")
)

type BookService struct {
	bookRepo        *repository.BookRepository
	reservationRepo *repository.ReservationRepository
}

func NewBookService(bookRepo *repository.BookRepository, reservationRepo *repository.ReservationRepository) *BookService {
	return &BookService{bookRepo: bookRepo, reservationRepo: reservationRepo}
}

func (s *BookService) Create(ctx context.Context, req dto.CreateBookRequest) (*dto.BookResponse, error) {
	existing, _ := s.bookRepo.GetByISBN(ctx, req.ISBN)
	if existing != nil {
		return nil, ErrISBNAlreadyExists
	}

	book := &model.Book{
		Title:           req.Title,
		Author:          req.Author,
		ISBN:            req.ISBN,
		Genre:           req.Genre,
		PublishedYear:   req.PublishedYear,
		TotalCopies:     req.TotalCopies,
		AvailableCopies: req.TotalCopies,
	}

	if err := s.bookRepo.Create(ctx, book); err != nil {
		return nil, err
	}

	return toBookResponse(book), nil
}

func (s *BookService) GetByID(ctx context.Context, id string) (*dto.BookResponse, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrBookNotFound
	}
	return toBookResponse(book), nil
}

func (s *BookService) Update(ctx context.Context, id string, req dto.UpdateBookRequest) (*dto.BookResponse, error) {
	book, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrBookNotFound
	}

	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.ISBN != nil {
		if *req.ISBN != book.ISBN {
			existing, _ := s.bookRepo.GetByISBN(ctx, *req.ISBN)
			if existing != nil {
				return nil, ErrISBNAlreadyExists
			}
		}
		book.ISBN = *req.ISBN
	}
	if req.Genre != nil {
		book.Genre = *req.Genre
	}
	if req.PublishedYear != nil {
		book.PublishedYear = *req.PublishedYear
	}
	if req.TotalCopies != nil {
		diff := *req.TotalCopies - book.TotalCopies
		book.TotalCopies = *req.TotalCopies
		book.AvailableCopies += diff
		if book.AvailableCopies < 0 {
			book.AvailableCopies = 0
		}
	}

	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, err
	}

	return toBookResponse(book), nil
}

func (s *BookService) Delete(ctx context.Context, id string) error {
	_, err := s.bookRepo.GetByID(ctx, id)
	if err != nil {
		return ErrBookNotFound
	}

	activeCount, err := s.reservationRepo.CountActiveByBookID(ctx, id)
	if err != nil {
		return err
	}
	if activeCount > 0 {
		return ErrBookHasReservations
	}

	return s.bookRepo.Delete(ctx, id)
}

func (s *BookService) Search(ctx context.Context, q dto.SearchBooksQuery) ([]dto.BookResponse, int, error) {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 || q.Limit > 50 {
		q.Limit = 10
	}

	books, totalCount, err := s.bookRepo.Search(ctx, q.Query, q.Genre, q.Available, q.Page, q.Limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.BookResponse
	for _, b := range books {
		responses = append(responses, *toBookResponse(&b))
	}

	return responses, totalCount, nil
}

func toBookResponse(b *model.Book) *dto.BookResponse {
	return &dto.BookResponse{
		ID:              b.ID,
		Title:           b.Title,
		Author:          b.Author,
		ISBN:            b.ISBN,
		Genre:           b.Genre,
		PublishedYear:   b.PublishedYear,
		Available:       b.Available(),
		TotalCopies:     b.TotalCopies,
		AvailableCopies: b.AvailableCopies,
	}
}
