package model

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	uuid "github.com/google/uuid"
)

// EmailTokenModel is an abstract model which can be used for objects from which
// we derive redirect emails (email confirmation, password reset and such)
type EmailToken struct {
	IDRecord
	Reference   uuid.UUID `bun:"type:uuid,notnull,default:uuid_generate_v4()"`
	EmailSent   bool      `bun:",notnull,default:false"`
	EmailSentAt *time.Time
	ExpiresAt   time.Time `bun:",notnull"`
}

type EmailTokenClaims struct {
	Username  string `json:"username"`
	Reference string `json:"reference"`
	jwt.StandardClaims
}

type Email struct {
	Recipient string
	Subject   string
	Template  string
}

// NewEmailToken creates new OauthEmailToken instance
func NewOauthEmailToken(expiresIn *time.Duration) *EmailToken {
	return &EmailToken{
		EmailSent: false,
		ExpiresAt: time.Now().UTC().Add(*expiresIn),
	}
}

// NewOauthEmailTokenClaims creates new NewOauthEmailTokenClaims instance
func NewOauthEmailTokenClaims(email string, emailToken *EmailToken) *EmailTokenClaims {
	return &EmailTokenClaims{
		Username:  email,
		Reference: emailToken.Reference.String(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: emailToken.ExpiresAt.Unix(),
		},
	}
}

// NewOauthEmail
func NewOauthEmail(email, subject, template string) *Email {
	return &Email{
		Recipient: email,
		Subject:   subject,
		Template:  template,
	}
}
