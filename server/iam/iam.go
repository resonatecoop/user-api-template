package iamserver

import (
	"context"

	"github.com/resonatecoop/user-api/model"
	"github.com/uptrace/bun"

	"github.com/rs/xid"

	"github.com/twitchtv/twirp"

	"github.com/resonatecoop/user-api/proto/iam"
	//"github.com/go-pg/pg/orm"
)

// New instantiates new IAM service
func New(db *bun.DB, tg TokenGenerator, udb UserDB, sec Securer) *Service {
	return &Service{db: db, tg: tg, udb: udb, sec: sec}
}

// Service represents IAM application service
type Service struct {
	db  *bun.DB
	tg  TokenGenerator
	udb UserDB
	sec Securer
}

// TokenGenerator generates new jwt token
type TokenGenerator interface {
	GenerateToken(*model.AuthUser) (string, error)
}

// UserDB represents user database interface
type UserDB interface {
	FindByAuth(context.Context, *bun.DB, string) (*model.User, error)
	FindByToken(context.Context, *bun.DB, string) (*model.User, error)
	UpdateLastLogin(context.Context, *bun.DB, *model.User) error
}

// Securer represents password securing service
type Securer interface {
	MatchesHash(string, string) bool
}

var (
	invalidUserPW = twirp.NewError(twirp.PermissionDenied, "invalid username or password")
	invalidToken  = twirp.NewError(twirp.PermissionDenied, "invalid token")
)

// Auth tries to authenticate user given username and password
func (s *Service) Auth(c context.Context, req *iam.AuthReq) (*iam.AuthResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// dbCtx := s.db.

	// WithContext(c)

	usr, err := s.udb.FindByAuth(c, s.db, req.Auth)
	if err != nil {
		return nil, invalidUserPW
	}

	if !s.sec.MatchesHash(usr.Password.String, req.Password) {
		return nil, invalidUserPW
	}

	token, err := s.tg.GenerateToken(&model.AuthUser{
		ID:       usr.ID,
		TenantID: usr.TenantID,
		Username: usr.Username,
		Email:    usr.Email,
		Role:     model.AccessRole(usr.RoleID),
	})

	if err != nil {
		return nil, err
	}

	uToken := xid.New().String()

	usr.UpdateLoginDetails(uToken)

	if err = s.udb.UpdateLastLogin(c, s.db, usr); err != nil {
		return nil, err
	}

	return &iam.AuthResp{
		Token:        token,
		RefreshToken: uToken,
	}, nil
}

// Refresh refreshes user's jwt token
func (s *Service) Refresh(c context.Context, req *iam.RefreshReq) (*iam.RefreshResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	usr, err := s.udb.FindByToken(c, s.db, req.Token)
	if err != nil {
		return nil, invalidToken
	}

	token, err := s.tg.GenerateToken(&model.AuthUser{
		ID:       usr.ID,
		TenantID: usr.TenantID,
		Username: usr.Username,
		Email:    usr.Email,
		Role:     model.AccessRole(usr.RoleID),
	})

	if err != nil {
		return nil, err
	}

	return &iam.RefreshResp{
		Token: token,
	}, nil
}
