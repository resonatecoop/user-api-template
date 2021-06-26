package server

import (
	"github.com/uptrace/bun"
)

// Server implements the UserService
type Server struct {
	db *bun.DB
}

// New creates an instance of our server
func New(db *bun.DB) *Server {
	return &Server{db: db}
}
