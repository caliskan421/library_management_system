package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/service"
)

type ReportHandler struct {
	reportService *service.ReportService
}

func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

// GET /api/reports
func (h *ReportHandler) GetReports(c *fiber.Ctx) error {
	var q dto.ReportQuery
	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Message: "Geçersiz sorgu parametreleri",
		})
	}

	result, err := h.reportService.Generate(c.Context(), q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: "Rapor oluşturulurken bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
