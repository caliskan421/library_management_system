package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/service"
)

type ReservationHandler struct {
	reservationService *service.ReservationService
}

func NewReservationHandler(reservationService *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{reservationService: reservationService}
}

// POST /api/reservations
func (h *ReservationHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var req dto.CreateReservationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz istek verisi",
		})
	}

	if req.BookID == "" || req.DueDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Kitap ID ve iade tarihi zorunludur",
		})
	}

	result, err := h.reservationService.Create(c.Context(), userID, req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrBookNotFound):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrBookNotAvailable):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrInvalidDueDate):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrMaxReservations):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Message: "Rezervasyon oluşturulurken bir hata oluştu",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// GET /api/reservations/:reservationid
func (h *ReservationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("reservationid")
	userID := c.Locals("userId").(string)
	userRole := c.Locals("userRole").(string)

	result, err := h.reservationService.GetByID(c.Context(), id, userID, userRole)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrReservationNotFound):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrNotOwnerOrAdmin):
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{Message: err.Error()})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Message: "Rezervasyon getirilirken bir hata oluştu",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// DELETE /api/reservations/:reservationid
func (h *ReservationHandler) Return(c *fiber.Ctx) error {
	id := c.Params("reservationid")
	userID := c.Locals("userId").(string)
	userRole := c.Locals("userRole").(string)

	err := h.reservationService.Return(c.Context(), id, userID, userRole)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrReservationNotFound):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrNotOwnerOrAdmin):
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{Message: err.Error()})
		case errors.Is(err, service.ErrAlreadyReturned):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Message: "İade işlemi sırasında bir hata oluştu",
			})
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GET /api/users/:userid/reservations
func (h *ReservationHandler) ListByUserID(c *fiber.Ctx) error {
	targetUserID := c.Params("userid")
	requestingUserID := c.Locals("userId").(string)
	userRole := c.Locals("userRole").(string)

	var q dto.UserReservationsQuery
	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz sorgu parametreleri",
		})
	}

	results, totalCount, err := h.reservationService.ListByUserID(c.Context(), targetUserID, requestingUserID, userRole, q)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotOwnerOrAdmin):
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{Message: err.Error()})
		default:
			if err.Error() == "kullanıcı bulunamadı" {
				return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Message: "Rezervasyonlar listelenirken bir hata oluştu",
			})
		}
	}

	if results == nil {
		results = []dto.ReservationResponse{}
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginatedResponse{
		Data:       results,
		Page:       q.Page,
		Limit:      q.Limit,
		TotalCount: totalCount,
	})
}
