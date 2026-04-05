package repository

import (
	"context"

	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/uptrace/bun"
)

type BookRepository struct {
	db *bun.DB
}

func NewBookRepository(db *bun.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(ctx context.Context, book *model.Book) error {
	_, err := r.db.NewInsert().Model(book).Exec(ctx)
	return err
}

func (r *BookRepository) GetByID(ctx context.Context, id string) (*model.Book, error) {
	book := new(model.Book)
	err := r.db.NewSelect().Model(book).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) GetByISBN(ctx context.Context, isbn string) (*model.Book, error) {
	book := new(model.Book)
	err := r.db.NewSelect().Model(book).Where("isbn = ?", isbn).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) Update(ctx context.Context, book *model.Book) error {
	_, err := r.db.NewUpdate().Model(book).WherePK().Exec(ctx)
	return err
}

func (r *BookRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.NewDelete().Model((*model.Book)(nil)).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *BookRepository) Search(ctx context.Context, query, genre string, available *bool, page, limit int) ([]model.Book, int, error) {
	var books []model.Book

	q := r.db.NewSelect().Model(&books)

	if query != "" {
		q = q.Where("(title ILIKE ? OR author ILIKE ? OR isbn ILIKE ?)",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}
	if genre != "" {
		q = q.Where("genre ILIKE ?", "%"+genre+"%")
	}
	if available != nil {
		if *available {
			q = q.Where("available_copies > 0")
		} else {
			q = q.Where("available_copies = 0")
		}
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = q.Limit(limit).Offset(offset).OrderExpr("created_at DESC").Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	return books, totalCount, nil
}

func (r *BookRepository) Count(ctx context.Context) (int, error) {
	return r.db.NewSelect().Model((*model.Book)(nil)).Count(ctx)
}
