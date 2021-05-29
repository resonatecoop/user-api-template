package model

import (
	"database/sql"
	"time"

	//	"github.com/RichardKnop/go-oauth2-server/util"
	// "github.com/RichardKnop/uuid"

	//"github.com/jinzhu/gorm"

	uuid "github.com/satori/go.uuid"
)

// OauthClient ...
type OauthClient struct {
	IDRecord
	Key                 string         `bun:"type:varchar(254);unique;not null"`
	Secret              string         `bun:"type:varchar(60);not null"`
	RedirectURI         sql.NullString `bun:"type:varchar(200)"`
	ApplicationName     sql.NullString `bun:"type:varchar(200)"`
	ApplicationHostname sql.NullString `bun:"type:varchar(200)"`
	ApplicationURL      sql.NullString `bun:"type:varchar(200)"`
}

// TableName specifies table name
func (c *OauthClient) TableName() string {
	return "oauth_clients"
}

// OauthScope ...
type OauthScope struct {
	IDRecord
	Scope       string `bun:"type:varchar(200);unique;not null"`
	Description sql.NullString
	IsDefault   bool `bun:"default:false"`
}

// TableName specifies table name
func (s *OauthScope) TableName() string {
	return "oauth_scopes"
}

// // OauthRole is a one of roles user can have (currently superuser or user)
// type OauthRole struct {
// 	TimestampModel
// 	ID   string `gorm:"primary_key" sql:"type:varchar(20)"`
// 	Name string `bun:"type:varchar(50);unique;not null"`
// }

// // TableName specifies table name
// func (r *OauthRole) TableName() string {
// 	return "oauth_roles"
// }

// OauthUser ...
// type OauthUser struct {
// 	RecordBasics
// 	RoleID         sql.NullString `bun:"type:varchar(20);index;not null"`
// 	Role           *OauthRole
// 	Username       string         `bun:"type:varchar(254);unique;not null"`
// 	Password       sql.NullString `bun:"type:varchar(60)"`
// 	EmailConfirmed bool           `bun:"default:false;not null"`
// }

// // TableName specifies table name
// func (u *OauthUser) TableName() string {
// 	return "oauth_users"
// }

// OauthRefreshToken ...
type OauthRefreshToken struct {
	IDRecord
	ClientID  sql.NullString `bun:"index;not null"`
	UserID    sql.NullString `bun:"index"`
	Client    *OauthClient
	User      *User
	Token     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	ExpiresAt time.Time `bun:"not null"`
	Scope     string    `bun:"type:varchar(200);not null"`
}

// TableName specifies table name
func (rt *OauthRefreshToken) TableName() string {
	return "oauth_refresh_tokens"
}

// OauthAccessToken ...
type OauthAccessToken struct {
	IDRecord
	ClientID  sql.NullString `bun:"index;not null"`
	UserID    sql.NullString `bun:"index"`
	Client    *OauthClient
	User      *User
	Token     string    `bun:"type:uuid,default:uuid_generate_v4()"`
	ExpiresAt time.Time `bun:"not null"`
	Scope     string    `bun:"type:varchar(200);not null"`
}

// TableName specifies table name
func (at *OauthAccessToken) TableName() string {
	return "oauth_access_tokens"
}

// OauthAuthorizationCode ...
type OauthAuthorizationCode struct {
	IDRecord
	ClientID    sql.NullString `bun:"index;not null"`
	UserID      sql.NullString `bun:"index;not null"`
	Client      *OauthClient
	User        *User
	Code        string         `bun:"type:uuid,default:uuid_generate_v4()"`
	RedirectURI sql.NullString `bun:"type:varchar(200)"`
	ExpiresAt   time.Time      `bun:"not null"`
	Scope       string         `bun:"type:varchar(200);not null"`
}

// TableName specifies table name
func (ac *OauthAuthorizationCode) TableName() string {
	return "oauth_authorization_codes"
}

// NewOauthRefreshToken creates new OauthRefreshToken instance
func NewOauthRefreshToken(client *OauthClient, user *User, expiresIn int, scope string) *OauthRefreshToken {
	refreshToken := &OauthRefreshToken{
		ClientID:  StringOrNull(client.ID.String()),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		refreshToken.UserID = StringOrNull(user.ID.String())
	}
	return refreshToken
}

// NewOauthAccessToken creates new OauthAccessToken instance
func NewOauthAccessToken(client *OauthClient, user *User, expiresIn int, scope string) *OauthAccessToken {
	accessToken := &OauthAccessToken{
		ClientID:  StringOrNull(client.ID.String()),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		accessToken.UserID = StringOrNull(user.ID.String())
	}
	return accessToken
}

// NewOauthAuthorizationCode creates new OauthAuthorizationCode instance
func NewOauthAuthorizationCode(client *OauthClient, user *User, expiresIn int, redirectURI, scope string) *OauthAuthorizationCode {
	return &OauthAuthorizationCode{
		ClientID:    StringOrNull(client.ID.String()),
		UserID:      StringOrNull(user.ID.String()),
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
