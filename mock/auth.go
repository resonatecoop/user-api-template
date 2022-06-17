package mock

import (
	"context"

	"github.com/resonatecoop/user-api-template/model"
)

// Auth mock
type Auth struct {
	GetUserFn func(context.Context) *model.AuthUser
}

// GetUser mock
func (s *Auth) GetUser(c context.Context) *model.AuthUser {
	return s.GetUserFn(c)
}
