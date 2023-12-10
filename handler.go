package main

import (
	"context"
	user_service "entry-rpc/kitex_gen/user_service"
	"entry-rpc/repository"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, userID int64) (resp *user_service.User, err error) {
	return repository.GetUserRepository().FindByID(ctx, userID)
}

// SaveUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) SaveUser(ctx context.Context, user *user_service.User) (err error) {
	return repository.GetUserRepository().Save(ctx, user)
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, userID int64) (err error) {
	return repository.GetUserRepository().DeleteByID(ctx, userID)
}
