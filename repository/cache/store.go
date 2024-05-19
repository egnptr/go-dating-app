package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/egnptr/dating-app/model"
)

func (cache *RedisCache) GetRelatedUserCacheLen(ctx context.Context, userID int64) (len int64, err error) {
	key := fmt.Sprintf("related_user:%d", userID)

	len, err = cache.Client.LLen(ctx, key).Result()
	if err != nil {
		err = errors.New("error fetching cached data")
		log.Println(err.Error())
		return
	}

	return
}

func (cache *RedisCache) GetRelatedUserCache(ctx context.Context, userID int64) (userRelationMap map[int64]int, err error) {
	userRelationMap = make(map[int64]int)
	key := fmt.Sprintf("related_user:%d", userID)

	cacheData, err := cache.Client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		err = errors.New("error fetching cached data")
		log.Println(err.Error())
		return
	}

	for _, str := range cacheData {
		var data model.UserRelation
		err = json.Unmarshal([]byte(str), &data)
		if err != nil {
			log.Println("error unmarshal json")
			return
		}

		userRelationMap[data.UserID] = data.SwipeStatus
	}

	return
}

func (cache *RedisCache) SetRelatedUserCache(ctx context.Context, userID int64, data model.UserRelation) (err error) {
	key := fmt.Sprintf("related_user:%d", userID)

	valueJson, err := json.Marshal(&data)
	if err != nil {
		log.Println("error marshal json")
		return
	}

	err = cache.Client.LPush(ctx, key, valueJson).Err()
	if err != nil {
		log.Println("error set cache: ", key)
		return
	}

	err = cache.Client.Expire(ctx, key, 24*time.Hour).Err()
	if err != nil {
		log.Println("error set cache expire: ", key)
		return
	}

	return
}
