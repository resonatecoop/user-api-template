package pgsql

import (
	"context"
	"strings"

	"github.com/resonatecoop/user-api/model"
	"github.com/uptrace/bun"
)

// NewUser returns a new User instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// FindByAuth finds user by either username or email
func (s *User) FindByAuth(ctx context.Context, db *bun.DB, auth string) (*model.User, error) {
	var user = new(model.User)

	lAuth := strings.ToLower(auth)
	err := db.NewSelect().
		Model(user).
		Where("lower(username) = ?", lAuth).
		Where("deleted_at is NULL").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByToken finds user by either username/email
func (s *User) FindByToken(ctx context.Context, db *bun.DB, token string) (*model.User, error) {
	var user = new(model.User)

	if dberr := db.NewSelect().
		Model(user).
		Where("token = ?", token).
		Where("deleted_at is NULL").
		Scan(ctx); dberr != nil {
		return nil, dberr
	}

	return user, nil
}

// UpdateLastLogin updates user's last login details
func (s *User) UpdateLastLogin(ctx context.Context, db *bun.DB, user *model.User) error {
	_, dberr := db.NewUpdate().
		Model(user).
		Column("last_login").
		WherePK().
		Exec(ctx)
	return dberr
}
