package cache

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/egnptr/dating-app/model"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var key = "related_user:1"

func TestGetRelatedUserCacheLen(t *testing.T) {
	type fields struct {
		redisClient *redis.Client
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes int64
		wantErr bool
	}{
		{
			name: "case success",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLLen(key).SetVal(4)
					return client
				}(),
			},
			args: args{
				id: 1,
			},
			wantRes: 4,
		},
		{
			name: "case error",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLLen(key).SetErr(errors.New("err"))
					return client
				}(),
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				Client: tt.fields.redisClient,
			}
			gotRes, gotErr := r.GetRelatedUserCacheLen(context.Background(), tt.args.id)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetRelatedUserCacheLen() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}

func TestGetRelatedUserCache(t *testing.T) {
	type fields struct {
		redisClient *redis.Client
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes map[int64]int
		wantErr bool
	}{
		{
			name: "case success",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLRange(key, 0, -1).SetVal([]string{
						`{"id": 2, "swipe_status": 1}`,
						`{"id": 3, "swipe_status": -1}`,
						`{"id": 4, "swipe_status": 1}`,
					})
					return client
				}(),
			},
			args: args{
				id: 1,
			},
			wantRes: map[int64]int{
				2: 1,
				3: -1,
				4: 1,
			},
		},
		{
			name: "case error",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLRange(key, 0, -1).SetErr(errors.New("err"))
					return client
				}(),
			},
			args: args{
				id: 1,
			},
			wantRes: map[int64]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				Client: tt.fields.redisClient,
			}
			gotRes, gotErr := r.GetRelatedUserCache(context.Background(), tt.args.id)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetRelatedUserCache() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}

func TestSetRelatedUserCache(t *testing.T) {
	data := model.UserRelation{
		UserID:      2,
		SwipeStatus: -1,
	}
	valueJson, _ := json.Marshal(&data)

	type fields struct {
		redisClient *redis.Client
	}
	type args struct {
		id   int64
		data model.UserRelation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case success",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLPush(key, valueJson).SetVal(1)
					mock.ExpectExpire(key, 24*time.Hour).SetVal(true)
					return client
				}(),
			},
			args: args{
				id:   1,
				data: data,
			},
		},
		{
			name: "case error LPush",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLPush(key, valueJson).SetErr(errors.New("err"))
					return client
				}(),
			},
			args: args{
				id:   1,
				data: data,
			},
			wantErr: true,
		},
		{
			name: "case error Expire",
			fields: fields{
				redisClient: func() *redis.Client {
					client, mock := redismock.NewClientMock()
					mock.ExpectLPush(key, valueJson).SetVal(1)
					mock.ExpectExpire(key, 24*time.Hour).SetErr(errors.New("err"))
					return client
				}(),
			},
			args: args{
				id:   1,
				data: data,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisCache{
				Client: tt.fields.redisClient,
			}
			gotErr := r.SetRelatedUserCache(context.Background(), tt.args.id, tt.args.data)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("SetRelatedUserCache() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}
