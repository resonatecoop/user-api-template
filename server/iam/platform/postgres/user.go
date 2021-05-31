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
	dberr := db.NewSelect().
		Model(user).
		//	Column("email", "username", "password").
		Where("lower(username) = ? or lower(email) = ?", lAuth, lAuth).
		Scan(ctx)

	if dberr != nil {
		return nil, dberr
	}

	return user, nil
}

// FindByToken finds user by either username or email
func (s *User) FindByToken(ctx context.Context, db *bun.DB, token string) (*model.User, error) {
	var user = new(model.User)

	if dberr := db.NewSelect().
		Model(user).
		Where("token = ?", token).
		// Where("deleted_at is null").
		Scan(ctx); dberr != nil {
		return nil, dberr
	}

	return user, nil
}

// UpdateLastLogin updates user's last login details
func (s *User) UpdateLastLogin(ctx context.Context, db *bun.DB, user *model.User) error {
	_, dberr := db.NewUpdate().
		Model(user).
		Column("last_login", "token").
		WherePK().
		Exec(ctx)
	return dberr
}
