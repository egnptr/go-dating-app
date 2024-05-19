package cache

import (
	"context"

	"github.com/egnptr/dating-app/model"
)

// go:generate moq -rm -out repo_mock.go . Repo
type Repo interface {
	GetRelatedUserCache(ctx context.Context, userID int64) (userRelationMap map[int64]int, err error)
	SetRelatedUserCache(ctx context.Context, userID int64, data model.UserRelation) (err error)
	GetRelatedUserCacheLen(ctx context.Context, userID int64) (len int64, err error)
}
