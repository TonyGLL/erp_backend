package db

import (
	"context"
)

type Querier interface {
	GetUser(ctx context.Context, id int32) (GetUserRow, error)
	GetUserPassword(ctx context.Context, username string) (string, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUserRow, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	CreatePassword(ctx context.Context, arg CreatePasswordParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	DeleteUser(ctx context.Context, arg DeleteUserParams) error
	GetUsersForDownload(ctx context.Context) ([]GetUserRow, error)
}

var _ Querier = (*Queries)(nil)
