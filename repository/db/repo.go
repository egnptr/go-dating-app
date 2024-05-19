package db

import (
	"context"

	"github.com/egnptr/dating-app/model"
)

// go:generate moq -rm -out repo_mock.go . Repo
type Repo interface {
	GetUser(ctx context.Context, username string) (*model.User, error)
	GetUserByID(ctx context.Context, userID int64) (*model.User, error)
	GetRelatedUser(ctx context.Context, id int64) ([]model.User, error)

	CreateUser(ctx context.Context, req model.User) (err error)
	UpdatePremiumStatus(ctx context.Context, req model.SubscribeRequest) (err error)
}
