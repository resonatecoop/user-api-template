package model

import (
	"database/sql"
	"time"

	//	"github.com/RichardKnop/go-oauth2-server/util"
	// "github.com/RichardKnop/uuid"

	//"github.com/jinzhu/gorm"

	uuid "github.com/google/uuid"
)

// OauthClient ...
type Client struct {
	IDRecord
	Key                 string         `bun:"type:varchar(254),unique,notnull"`
	Secret              string         `bun:"type:varchar(60),notnull"`
	RedirectURI         sql.NullString `bun:"type:varchar(200)"`
	ApplicationName     sql.NullString `bun:"type:varchar(200)"`
	ApplicationHostname sql.NullString `bun:"type:varchar(200)"`
	ApplicationURL      sql.NullString `bun:"type:varchar(200)"`
}

// // TableName specifies table name
// func (c *OauthClient) TableName() string {
// 	return "oauth_clients"
// }

// OauthScope ...
type Scope struct {
	ID          int32  `bun:"type:,unique"`
	Name        string `bun:"type:varchar(50),unique,notnull"`
	Description string `bun:"type:varchar(200)"`
	IsDefault   bool   `bun:"default:false"`
}

// // TableName specifies table name
// func (s *OauthScope) TableName() string {
// 	return "oauth_scopes"
// }

// // Role is a one of roles user can have (currently superuser or user)
// type Role struct {
// 	ID          int8   `bun:"primary_key" sql:"type:varchar(20)"`
// 	Name        string `bun:"type:varchar(50),unique,notnull"`
// 	Description string `bun:"type:varchar(200),notnull"`
// }

// // TableName specifies table name
// func (r *OauthRole) TableName() string {
// 	return "oauth_roles"
// }

// OauthUser ...
// type OauthUser struct {
// 	RecordBasics
// 	RoleID         sql.NullString `bun:"type:varchar(20);index,notnull"`
// 	Role           *OauthRole
// 	Username       string         `bun:"type:varchar(254),unique,notnull"`
// 	Password       sql.NullString `bun:"type:varchar(60)"`
// 	EmailConfirmed bool           `bun:"default:false,notnull"`
// }

// // TableName specifies table name
// func (u *OauthUser) TableName() string {
// 	return "oauth_users"
// }

// OauthRefreshToken ...
type RefreshToken struct {
	IDRecord
	ClientID  uuid.UUID `bun:"type:uuid,notnull"`
	UserID    uuid.UUID `bun:"type:uuid"`
	Client    *Client   `bun:"rel:has-one"`
	User      *User     `bun:"rel:has-one"`
	Token     string    `bun:"type:varchar(40),unique,notnull"`
	ExpiresAt time.Time `bun:",notnull,default:now()"`
	Scope     string    `bun:"type:varchar(200),notnull"`
}

// // TableName specifies table name
// func (rt *OauthRefreshToken) TableName() string {
// 	return "oauth_refresh_tokens"
// }

// OauthAccessToken ...
type AccessToken struct {
	IDRecord
	ClientID  uuid.UUID `bun:"type:uuid,notnull"`
	UserID    uuid.UUID `bun:"type:uuid"`
	Client    *Client   `bun:"rel:has-one"`
	User      *User     `bun:"rel:has-one"`
	Token     string    `bun:"type:varchar(40),unique,notnull"`
	ExpiresAt time.Time `bun:",notnull"`
	Scope     string    `bun:"type:varchar(200),notnull"`
}

// // TableName specifies table name
// func (at *OauthAccessToken) TableName() string {
// 	return "oauth_access_tokens"
// }

// OauthAuthorizationCode ...
type AuthorizationCode struct {
	IDRecord
	ClientID    uuid.UUID      `bun:"type:uuid,notnull"`
	UserID      uuid.UUID      `bun:"type:uuid,notnull"`
	Client      *Client        `bun:"rel:has-one"`
	User        *User          `bun:"rel:has-one"`
	Code        string         `bun:"type:varchar(40),unique,notnull"`
	RedirectURI sql.NullString `bun:"type:varchar(200)"`
	ExpiresAt   time.Time      `bun:",notnull"`
	Scope       string         `bun:"type:varchar(200),notnull"`
}

// // TableName specifies table name
// func (ac *OauthAuthorizationCode) TableName() string {
// 	return "oauth_authorization_codes"
// }

// NewOauthRefreshToken creates new OauthRefreshToken instance
func NewOauthRefreshToken(client *Client, user *User, expiresIn int, scope string) *RefreshToken {
	refreshToken := &RefreshToken{
		IDRecord:  IDRecord{CreatedAt: time.Now().UTC()},
		ClientID:  client.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		refreshToken.UserID = user.ID
	}
	return refreshToken
}

// NewOauthAccessToken creates new OauthAccessToken instance
func NewOauthAccessToken(client *Client, user *User, expiresIn int, scope string) *AccessToken {
	accessToken := &AccessToken{
		IDRecord:  IDRecord{CreatedAt: time.Now().UTC()},
		ClientID:  client.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		accessToken.UserID = user.ID
	}
	return accessToken
}

// NewOauthAuthorizationCode creates new OauthAuthorizationCode instance
func NewOauthAuthorizationCode(client *Client, user *User, expiresIn int, redirectURI, scope string) *AuthorizationCode {
	return &AuthorizationCode{
		IDRecord:    IDRecord{CreatedAt: time.Now().UTC()},
		ClientID:    client.ID,
		UserID:      user.ID,
		Code:        uuid.New().String(),
		ExpiresAt:   time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		RedirectURI: StringOrNull(redirectURI),
		Scope:       scope,
	}
}

// OauthAuthorizationCodePreload sets up Gorm preloads for an auth code object
// func OauthAuthorizationCodePreload(db *bun.DB) *bun.DB {
// 	return OauthAuthorizationCodePreloadWithPrefix(db, "")
// }

// // OauthAuthorizationCodePreloadWithPrefix sets up Gorm preloads for an auth code object,
// // and prefixes with prefix for nested objects
// func OauthAuthorizationCodePreloadWithPrefix(db *bun.DB, prefix string) *bun.DB {
// 	return db.
// 		Preload(prefix + "Client").Preload(prefix + "User")
// }

// // OauthAccessTokenPreload sets up Gorm preloads for an access token object
// func OauthAccessTokenPreload(db *bun.DB) *bun.DB {
// 	return OauthAccessTokenPreloadWithPrefix(db, "")
// }

// OauthAccessTokenPreload sets up Gorm preloads for an access token object
// func OauthAccessTokenPreload(db *bun.DB) *bun.DB {
// 	return db.Relation("Client").Relation("User")
// }

// OauthAccessTokenPreloadWithPrefix sets up Gorm preloads for an access token object,
// and prefixes with prefix for nested objects
// func OauthAccessTokenPreloadWithPrefix(db *bun.DB, prefix string) *bun.DB {
// 	return db.
// 		Preload(prefix + "Client").Preload(prefix + "User")
// }

// // OauthRefreshTokenPreload sets up Gorm preloads for a refresh token object
// func OauthRefreshTokenPreload(db *bun.DB) *bun.DB {
// 	return OauthRefreshTokenPreloadWithPrefix(db, "")
// }

// // OauthRefreshTokenPreloadWithPrefix sets up Gorm preloads for a refresh token object,
// // and prefixes with prefix for nested objects
// func OauthRefreshTokenPreloadWithPrefix(db *bun.DB, prefix string) *bun.DB {
// 	return db.
// 		Preload(prefix + "Client").Preload(prefix + "User")
// }

// StringOrNull returns properly configured sql.NullString
func StringOrNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}
