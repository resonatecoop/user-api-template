package server

import (
	"context"
	"sync"

	"github.com/gofrs/uuid"

	pbUser "github.com/merefield/grpc-user-api/proto"
)

// Backend implements the protobuf interface
type Backend struct {
	mu    *sync.RWMutex
	users []*pbUser.User
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

// AddUser gets a user to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, _ *pbUser.AddUserRequest) (*pbUser.User, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &pbUser.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return user, nil
}

// GetUser Gets a user to the in-memory store.
func (b *Backend) GetUser(ctx context.Context, _ *pbUser.GetUserRequest) (*pbUser.User, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &pbUser.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return user, nil
}

// UpdateUser gets a user to the in-memory store.
func (b *Backend) UpdateUser(ctx context.Context, _ *pbUser.UpdateUserRequest) (*pbUser.User, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &pbUser.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return user, nil
}

// ListUsers lists all users in the store.
func (b *Backend) ListUsers(_ *pbUser.ListUsersRequest, srv pbUser.UserService_ListUsersServer) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, user := range b.users {
		err := srv.Send(user)
		if err != nil {
			return err
		}
	}

	return nil
}
