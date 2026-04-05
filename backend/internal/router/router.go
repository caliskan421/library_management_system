package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetalicaliskan/libranet/internal/handler"
	"github.com/muhammetalicaliskan/libranet/internal/middleware"
)

type Handlers struct {
	Auth        *handler.AuthHandler
	Book        *handler.BookHandler
	Reservation *handler.ReservationHandler
	Report      *handler.ReportHandler
}

func Setup(app *fiber.App, h Handlers, jwtSecret string) {
	api := app.Group("/api")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.Login)

	// Protected routes
	protected := api.Group("", middleware.Authenticate(jwtSecret))

	// Book routes
	books := protected.Group("/books")
	books.Get("/", h.Book.Search)
	books.Get("/:bookid", h.Book.GetByID)
	books.Post("/", middleware.Authorize("admin"), h.Book.Create)
	books.Put("/:bookid", middleware.Authorize("admin"), h.Book.Update)
	books.Delete("/:bookid", middleware.Authorize("admin"), h.Book.Delete)

	// Reservation routes
	reservations := protected.Group("/reservations")
	reservations.Post("/", h.Reservation.Create)
	reservations.Get("/:reservationid", h.Reservation.GetByID)
	reservations.Delete("/:reservationid", h.Reservation.Return)

	// User reservations
	protected.Get("/users/:userid/reservations", h.Reservation.ListByUserID)

	// Report routes (admin only)
	reports := protected.Group("/reports", middleware.Authorize("admin"))
	reports.Get("/", h.Report.GetReports)
}
