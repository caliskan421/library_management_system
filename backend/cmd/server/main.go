package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/muhammetalicaliskan/libranet/internal/config"
	"github.com/muhammetalicaliskan/libranet/internal/database"
	"github.com/muhammetalicaliskan/libranet/internal/handler"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
	"github.com/muhammetalicaliskan/libranet/internal/router"
	"github.com/muhammetalicaliskan/libranet/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg.DB)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	if err := database.Migrate(context.Background(), db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database connected and migrated successfully")

	// Repositories
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	reservationRepo := repository.NewReservationRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, cfg)
	bookService := service.NewBookService(bookRepo, reservationRepo)
	reservationService := service.NewReservationService(reservationRepo, bookRepo, userRepo)
	reportService := service.NewReportService(reservationRepo, bookRepo, userRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	bookHandler := handler.NewBookHandler(bookService)
	reservationHandler := handler.NewReservationHandler(reservationService)
	reportHandler := handler.NewReportHandler(reportService)

	// Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"message": err.Error()})
		},
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Temporary seed endpoint - promote user to admin by secret key
	app.Post("/seed/admin", func(c *fiber.Ctx) error {
		var body struct {
			Email string `json:"email"`
			Key   string `json:"key"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "invalid request"})
		}
		if body.Key != cfg.JWT.Secret {
			return c.Status(403).JSON(fiber.Map{"message": "invalid key"})
		}
		_, err := db.ExecContext(c.Context(), "UPDATE users SET role='admin' WHERE email=?", body.Email)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "user promoted to admin"})
	})

	// Routes
	router.Setup(app, router.Handlers{
		Auth:        authHandler,
		Book:        bookHandler,
		Reservation: reservationHandler,
		Report:      reportHandler,
	}, cfg.JWT.Secret)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.App.Port)
		log.Printf("LibraNet server starting on %s (env: %s)", addr, cfg.App.Env)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server stopped gracefully")
}
