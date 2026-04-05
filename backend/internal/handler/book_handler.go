package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/service"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

// POST /api/books
func (h *BookHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz istek verisi",
		})
	}

	if req.Title == "" || req.Author == "" || req.ISBN == "" || req.TotalCopies < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Başlık, yazar, ISBN ve kopya sayısı zorunludur",
		})
	}

	result, err := h.bookService.Create(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrISBNAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kitap eklenirken bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// GET /api/books/:bookid
func (h *BookHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("bookid")

	result, err := h.bookService.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kitap getirilirken bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// PUT /api/books/:bookid
func (h *BookHandler) Update(c *fiber.Ctx) error {
	id := c.Params("bookid")

	var req dto.UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz istek verisi",
		})
	}

	result, err := h.bookService.Update(c.Context(), id, req)
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		if errors.Is(err, service.ErrISBNAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kitap güncellenirken bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// DELETE /api/books/:bookid
func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("bookid")

	err := h.bookService.Delete(c.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrBookNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		if errors.Is(err, service.ErrBookHasReservations) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kitap silinirken bir hata oluştu",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GET /api/books
func (h *BookHandler) Search(c *fiber.Ctx) error {
	var q dto.SearchBooksQuery
	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz sorgu parametreleri",
		})
	}

	results, totalCount, err := h.bookService.Search(c.Context(), q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Kitaplar aranırken bir hata oluştu",
		})
	}

	if results == nil {
		results = []dto.BookResponse{}
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginatedResponse{
		Data:       results,
		Page:       q.Page,
		Limit:      q.Limit,
		TotalCount: totalCount,
	})
}
