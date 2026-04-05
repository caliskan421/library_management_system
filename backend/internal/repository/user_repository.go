package repository

import (
	"context"

	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Count(ctx context.Context) (int, error) {
	return r.db.NewSelect().Model((*model.User)(nil)).Count(ctx)
}
