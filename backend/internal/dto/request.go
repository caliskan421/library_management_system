package dto

// Auth
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Book
type CreateBookRequest struct {
	Title         string `json:"title" validate:"required,min=1,max=200"`
	Author        string `json:"author" validate:"required,min=2,max=100"`
	ISBN          string `json:"isbn" validate:"required"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"publishedYear"`
	TotalCopies   int    `json:"totalCopies" validate:"required,min=1"`
}

type UpdateBookRequest struct {
	Title         *string `json:"title"`
	Author        *string `json:"author"`
	ISBN          *string `json:"isbn"`
	Genre         *string `json:"genre"`
	PublishedYear *int    `json:"publishedYear"`
	TotalCopies   *int    `json:"totalCopies"`
}

type SearchBooksQuery struct {
	Query     string `query:"query"`
	Genre     string `query:"genre"`
	Available *bool  `query:"available"`
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
}

// Reservation
type CreateReservationRequest struct {
	BookID  string `json:"bookId" validate:"required"`
	DueDate string `json:"dueDate" validate:"required"`
}

// Report
type ReportQuery struct {
	Type string `query:"type"`
	From string `query:"from"`
	To   string `query:"to"`
}

// User Reservations
type UserReservationsQuery struct {
	Status string `query:"status"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}
