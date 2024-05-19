package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/egnptr/dating-app/model"
	"github.com/egnptr/dating-app/pkg/util"
	"github.com/egnptr/dating-app/repository/cache"
	"github.com/egnptr/dating-app/repository/db"
)

type usecase struct {
	RepoDB    db.Repo
	RepoCache cache.Repo
}

func (s *usecase) CreateUser(ctx context.Context, req model.User) (err error) {
	var hashedPassword string

	hashedPassword, err = util.HashPassword(req.Password)
	if err != nil {
		log.Println("error hashing password")
		return
	}
	req.Password = hashedPassword

	err = s.RepoDB.CreateUser(ctx, req)
	if err != nil {
		log.Println("error when creating new user in db")
	}

	return
}

func (s *usecase) Login(ctx context.Context, req model.LoginRequest) (err error) {
	user, err := s.RepoDB.GetUser(ctx, req.Username)
	if err != nil {
		log.Println("error when fetching user from db")
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		err = model.UnauthorizedErr
		log.Println("error unauthorized login")
	}

	return
}

func (s *usecase) UpdateSubscription(ctx context.Context, req model.SubscribeRequest) (err error) {
	err = s.RepoDB.UpdatePremiumStatus(ctx, req)
	if err != nil {
		log.Println("error when updating subscription status from db")
	}

	return
}

func (s *usecase) GetProfiles(ctx context.Context, req model.GetRelatedUserRequest) (filteredUser []model.User, err error) {
	users, err := s.RepoDB.GetRelatedUser(ctx, req.UserID)
	if err != nil {
		log.Println("error when fetching related users from db")
		return
	}

	if len(users) == 0 {
		err = errors.New("error empty users")
		log.Println(err.Error())
		return
	}

	userRelationMap, err := s.RepoCache.GetRelatedUserCache(ctx, req.UserID)
	if err != nil {
		log.Println("error when fetching related users from cache")
		return
	}

	for _, user := range users {
		_, exist := userRelationMap[user.UserID]
		if exist {
			continue
		}

		filteredUser = append(filteredUser, user)
	}

	return
}

func (s *usecase) Swipe(ctx context.Context, req model.SwipeRequest) (err error) {
	user, err := s.RepoDB.GetUserByID(ctx, req.UserID)
	if err != nil {
		log.Println("error when fetching user from db")
		return
	}

	// Limit number of swipes based on subscription status
	if !user.IsPremium {
		limit, errCache := s.RepoCache.GetRelatedUserCacheLen(ctx, req.UserID)
		if errCache != nil {
			err = errCache
			log.Println("error when fetching related users length from cache")
			return
		}

		if limit > 10 {
			err = errors.New("error reach limit of swipe")
			return
		}
	}

	err = s.RepoCache.SetRelatedUserCache(ctx, req.UserID, model.UserRelation{
		UserID:      req.SwipedUserID,
		SwipeStatus: req.SwipeStatus,
	})
	if err != nil {
		log.Println("error when setting related users to cache")
		return
	}

	return
}
