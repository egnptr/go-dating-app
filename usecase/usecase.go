package usecase

import (
	"context"

	"github.com/egnptr/dating-app/model"
	"github.com/egnptr/dating-app/repository/cache"
	"github.com/egnptr/dating-app/repository/db"
)

// go:generate moq -rm -out usecase_mock.go . Usecases
type Usecases interface {
	CreateUser(ctx context.Context, req model.User) error
	Login(ctx context.Context, req model.LoginRequest) error
	UpdateSubscription(ctx context.Context, req model.SubscribeRequest) (err error)
	GetProfiles(ctx context.Context, req model.GetRelatedUserRequest) (filteredUser []model.User, err error)
	Swipe(ctx context.Context, req model.SwipeRequest) (err error)
}

type usecase struct {
	RepoDB    db.Repo
	RepoCache cache.Repo
}

func NewUsecase(db db.Repo, cache cache.Repo) Usecases {
	return &usecase{
		RepoDB:    db,
		RepoCache: cache,
	}
}
