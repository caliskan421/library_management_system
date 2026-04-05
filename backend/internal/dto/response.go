package dto

import "time"

// Auth
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

// Book
type BookResponse struct {
	ID              string `json:"_id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Genre           string `json:"genre"`
	PublishedYear   int    `json:"publishedYear"`
	Available       bool   `json:"available"`
	TotalCopies     int    `json:"totalCopies"`
	AvailableCopies int    `json:"availableCopies"`
}

// Reservation
type ReservationResponse struct {
	ID         string     `json:"_id"`
	UserID     string     `json:"userId"`
	BookID     string     `json:"bookId"`
	Status     string     `json:"status"`
	ReservedAt time.Time  `json:"reservedAt"`
	DueDate    time.Time  `json:"dueDate"`
	ReturnedAt *time.Time `json:"returnedAt"`
}

// Report
type ReportResponse struct {
	GeneratedAt  time.Time                 `json:"generatedAt"`
	Type         string                    `json:"type"`
	Summary      ReportSummary             `json:"summary"`
	TopBooks     []TopBookReport           `json:"topBooks,omitempty"`
	Reservations []ReservationDetailReport `json:"reservations,omitempty"`
}

type ReportSummary struct {
	Total    int `json:"total"`
	Active   int `json:"active"`
	Returned int `json:"returned"`
}

type TopBookReport struct {
	BookID           string `json:"bookId"`
	Title            string `json:"title"`
	ReservationCount int    `json:"reservationCount"`
}

type ReservationDetailReport struct {
	BookTitle  string `json:"bookTitle"`
	UserName   string `json:"userName"`
	UserEmail  string `json:"userEmail"`
	Status     string `json:"status"`
	ReservedAt string `json:"reservedAt"`
	DueDate    string `json:"dueDate"`
}

// Error
type ErrorResponse struct {
	Message string `json:"message"`
}

// Pagination
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalCount int         `json:"totalCount"`
}
