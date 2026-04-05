package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books,alias:b"`

	ID              string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"_id"`
	Title           string    `bun:"title,notnull" json:"title"`
	Author          string    `bun:"author,notnull" json:"author"`
	ISBN            string    `bun:"isbn,notnull,unique" json:"isbn"`
	Genre           string    `bun:"genre" json:"genre"`
	PublishedYear   int       `bun:"published_year" json:"publishedYear"`
	TotalCopies     int       `bun:"total_copies,notnull,default:1" json:"totalCopies"`
	AvailableCopies int       `bun:"available_copies,notnull,default:1" json:"availableCopies"`
	CreatedAt       time.Time `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updatedAt"`
}

func (b *Book) Available() bool {
	return b.AvailableCopies > 0
}
