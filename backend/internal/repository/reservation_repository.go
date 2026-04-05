package repository

import (
	"context"
	"time"

	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/uptrace/bun"
)

type ReservationRepository struct {
	db *bun.DB
}

func NewReservationRepository(db *bun.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(ctx context.Context, reservation *model.Reservation) error {
	_, err := r.db.NewInsert().Model(reservation).Exec(ctx)
	return err
}

func (r *ReservationRepository) GetByID(ctx context.Context, id string) (*model.Reservation, error) {
	reservation := new(model.Reservation)
	err := r.db.NewSelect().Model(reservation).Where("r.id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (r *ReservationRepository) Update(ctx context.Context, reservation *model.Reservation) error {
	_, err := r.db.NewUpdate().Model(reservation).WherePK().Exec(ctx)
	return err
}

func (r *ReservationRepository) ListByUserID(ctx context.Context, userID, status string, page, limit int) ([]model.Reservation, int, error) {
	var reservations []model.Reservation

	q := r.db.NewSelect().Model(&reservations).Where("user_id = ?", userID)

	if status != "" && status != "all" {
		q = q.Where("status = ?", status)
	}

	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = q.Limit(limit).Offset(offset).OrderExpr("reserved_at DESC").Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	return reservations, totalCount, nil
}

func (r *ReservationRepository) CountActiveByBookID(ctx context.Context, bookID string) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("book_id = ?", bookID).
		Where("status = ?", model.StatusActive).
		Count(ctx)
}

func (r *ReservationRepository) CountActiveByUserID(ctx context.Context, userID string) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("user_id = ?", userID).
		Where("status = ?", model.StatusActive).
		Count(ctx)
}

func (r *ReservationRepository) TotalCount(ctx context.Context) (int, error) {
	return r.db.NewSelect().Model((*model.Reservation)(nil)).Count(ctx)
}

func (r *ReservationRepository) CountByStatus(ctx context.Context, status model.ReservationStatus) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("status = ?", status).
		Count(ctx)
}

func (r *ReservationRepository) CountByDateRange(ctx context.Context, from, to time.Time) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("reserved_at >= ?", from).
		Where("reserved_at <= ?", to).
		Count(ctx)
}

func (r *ReservationRepository) CountActiveByDateRange(ctx context.Context, from, to time.Time) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("reserved_at >= ?", from).
		Where("reserved_at <= ?", to).
		Where("status = ?", model.StatusActive).
		Count(ctx)
}

func (r *ReservationRepository) CountReturnedByDateRange(ctx context.Context, from, to time.Time) (int, error) {
	return r.db.NewSelect().
		Model((*model.Reservation)(nil)).
		Where("reserved_at >= ?", from).
		Where("reserved_at <= ?", to).
		Where("status = ?", model.StatusReturned).
		Count(ctx)
}

type TopBook struct {
	BookID           string `bun:"book_id"`
	Title            string `bun:"title"`
	ReservationCount int    `bun:"reservation_count"`
}

func (r *ReservationRepository) TopReservedBooks(ctx context.Context, limit int) ([]TopBook, error) {
	var results []TopBook
	err := r.db.NewSelect().
		TableExpr("reservations AS r").
		Join("JOIN books AS b ON b.id = r.book_id").
		ColumnExpr("r.book_id").
		ColumnExpr("b.title").
		ColumnExpr("COUNT(*) AS reservation_count").
		GroupExpr("r.book_id, b.title").
		OrderExpr("reservation_count DESC").
		Limit(limit).
		Scan(ctx, &results)
	return results, err
}
