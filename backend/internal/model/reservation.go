package model

import (
	"time"

	"github.com/uptrace/bun"
)

type ReservationStatus string

const (
	StatusActive   ReservationStatus = "active"
	StatusReturned ReservationStatus = "returned"
)

type Reservation struct {
	bun.BaseModel `bun:"table:reservations,alias:r"`

	ID         string            `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"_id"`
	UserID     string            `bun:"user_id,notnull,type:uuid" json:"userId"`
	BookID     string            `bun:"book_id,notnull,type:uuid" json:"bookId"`
	Status     ReservationStatus `bun:"status,notnull,default:'active'" json:"status"`
	ReservedAt time.Time         `bun:"reserved_at,notnull,default:current_timestamp" json:"reservedAt"`
	DueDate    time.Time         `bun:"due_date,notnull" json:"dueDate"`
	ReturnedAt *time.Time        `bun:"returned_at" json:"returnedAt"`
	CreatedAt  time.Time         `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt  time.Time         `bun:"updated_at,notnull,default:current_timestamp" json:"updatedAt"`

	User *User `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	Book *Book `bun:"rel:belongs-to,join:book_id=id" json:"book,omitempty"`
}
