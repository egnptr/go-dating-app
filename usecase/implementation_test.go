package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/egnptr/dating-app/model"
	"github.com/egnptr/dating-app/pkg/util"
	"github.com/egnptr/dating-app/repository/cache"
	"github.com/egnptr/dating-app/repository/db"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	type fields struct {
		repoDB    db.Repo
		repoCache cache.Repo
	}
	type args struct {
		req model.User
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
				repoDB: &db.RepoMock{
					CreateUserFunc: func(ctx context.Context, req model.User) error {
						return nil
					},
				},
			},
			args: args{
				req: model.User{
					Username: "test",
					Password: "password",
					FullName: "full name",
					Email:    "test@mail.com",
				},
			},
		},
		{
			name: "case error db",
			fields: fields{
				repoDB: &db.RepoMock{
					CreateUserFunc: func(ctx context.Context, req model.User) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				req: model.User{
					Username: "test",
					Password: "password",
					FullName: "full name",
					Email:    "test@mail.com",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				RepoDB:    tt.fields.repoDB,
				RepoCache: tt.fields.repoCache,
			}
			gotErr := u.CreateUser(context.Background(), tt.args.req)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}

func TestLogin(t *testing.T) {
	password := "testpass"
	hashedPassword, _ := util.HashPassword(password)

	type fields struct {
		repoDB    db.Repo
		repoCache cache.Repo
	}
	type args struct {
		req model.LoginRequest
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
				repoDB: &db.RepoMock{
					GetUserFunc: func(ctx context.Context, username string) (*model.User, error) {
						return &model.User{
							Username: "test",
							Password: hashedPassword,
						}, nil
					},
				},
			},
			args: args{
				req: model.LoginRequest{
					Username: "test",
					Password: password,
				},
			},
		},
		{
			name: "case error db",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserFunc: func(ctx context.Context, username string) (*model.User, error) {
						return &model.User{}, errors.New("err")
					},
				},
			},
			args: args{
				req: model.LoginRequest{
					Username: "test",
					Password: password,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				RepoDB:    tt.fields.repoDB,
				RepoCache: tt.fields.repoCache,
			}
			gotErr := u.Login(context.Background(), tt.args.req)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateSubscription(t *testing.T) {
	type fields struct {
		repoDB    db.Repo
		repoCache cache.Repo
	}
	type args struct {
		req model.SubscribeRequest
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
				repoDB: &db.RepoMock{
					UpdatePremiumStatusFunc: func(ctx context.Context, req model.SubscribeRequest) error {
						return nil
					},
				},
			},
			args: args{
				req: model.SubscribeRequest{
					UserID:    1,
					Subscribe: true,
				},
			},
		},
		{
			name: "case erorr db",
			fields: fields{
				repoDB: &db.RepoMock{
					UpdatePremiumStatusFunc: func(ctx context.Context, req model.SubscribeRequest) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				req: model.SubscribeRequest{
					UserID:    1,
					Subscribe: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				RepoDB:    tt.fields.repoDB,
				RepoCache: tt.fields.repoCache,
			}
			gotErr := u.UpdateSubscription(context.Background(), tt.args.req)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("UpdateSubscription() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}

func TestGetProfiles(t *testing.T) {
	type fields struct {
		repoDB    db.Repo
		repoCache cache.Repo
	}
	type args struct {
		req model.GetRelatedUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []model.User
		wantErr bool
	}{
		{
			name: "case success",
			fields: fields{
				repoDB: &db.RepoMock{
					GetRelatedUserFunc: func(ctx context.Context, id int64) ([]model.User, error) {
						return []model.User{
							{
								UserID: 2,
							},
							{
								UserID: 3,
							},
							{
								UserID: 4,
							},
							{
								UserID: 5,
							},
							{
								UserID: 6,
							},
							{
								UserID: 7,
							},
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					GetRelatedUserCacheFunc: func(ctx context.Context, userID int64) (map[int64]int, error) {
						return map[int64]int{
							4: 0,
							7: 1,
						}, nil
					},
				},
			},
			args: args{
				req: model.GetRelatedUserRequest{
					UserID: 1,
				},
			},
			wantRes: []model.User{
				{
					UserID: 2,
				},
				{
					UserID: 3,
				},
				{
					UserID: 5,
				},
				{
					UserID: 6,
				},
			},
		},
		{
			name: "case error db",
			fields: fields{
				repoDB: &db.RepoMock{
					GetRelatedUserFunc: func(ctx context.Context, id int64) ([]model.User, error) {
						return []model.User{}, errors.New("err")
					},
				},
			},
			args: args{
				req: model.GetRelatedUserRequest{
					UserID: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "case error empty db",
			fields: fields{
				repoDB: &db.RepoMock{
					GetRelatedUserFunc: func(ctx context.Context, id int64) ([]model.User, error) {
						return []model.User{}, nil
					},
				},
			},
			args: args{
				req: model.GetRelatedUserRequest{
					UserID: 1,
				},
			},
			wantErr: true,
		},
		{
			name: "case error cache",
			fields: fields{
				repoDB: &db.RepoMock{
					GetRelatedUserFunc: func(ctx context.Context, id int64) ([]model.User, error) {
						return []model.User{
							{
								UserID: 2,
							},
							{
								UserID: 3,
							},
							{
								UserID: 4,
							},
							{
								UserID: 5,
							},
							{
								UserID: 6,
							},
							{
								UserID: 7,
							},
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					GetRelatedUserCacheFunc: func(ctx context.Context, userID int64) (map[int64]int, error) {
						return map[int64]int{}, errors.New("err")
					},
				},
			},
			args: args{
				req: model.GetRelatedUserRequest{
					UserID: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				RepoDB:    tt.fields.repoDB,
				RepoCache: tt.fields.repoCache,
			}
			gotRes, gotErr := u.GetProfiles(context.Background(), tt.args.req)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetProfiles() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}

func TestSwipe(t *testing.T) {
	type fields struct {
		repoDB    db.Repo
		repoCache cache.Repo
	}
	type args struct {
		req model.SwipeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case success not premium",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserByIDFunc: func(ctx context.Context, userID int64) (*model.User, error) {
						return &model.User{
							UserID:    1,
							IsPremium: false,
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					GetRelatedUserCacheLenFunc: func(ctx context.Context, userID int64) (int64, error) {
						return 6, nil
					},
					SetRelatedUserCacheFunc: func(ctx context.Context, userID int64, data model.UserRelation) error {
						return nil
					},
				},
			},
			args: args{
				req: model.SwipeRequest{
					UserID:       1,
					SwipedUserID: 2,
					SwipeStatus:  -1,
				},
			},
		},
		{
			name: "case success premium",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserByIDFunc: func(ctx context.Context, userID int64) (*model.User, error) {
						return &model.User{
							UserID:    1,
							IsPremium: true,
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					SetRelatedUserCacheFunc: func(ctx context.Context, userID int64, data model.UserRelation) error {
						return nil
					},
				},
			},
			args: args{
				req: model.SwipeRequest{
					UserID:       1,
					SwipedUserID: 2,
					SwipeStatus:  -1,
				},
			},
		},
		{
			name: "case error db",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserByIDFunc: func(ctx context.Context, userID int64) (*model.User, error) {
						return &model.User{}, errors.New("err")
					},
				},
			},
			args: args{
				req: model.SwipeRequest{
					UserID:       1,
					SwipedUserID: 2,
					SwipeStatus:  -1,
				},
			},
			wantErr: true,
		},
		{
			name: "case error cache len",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserByIDFunc: func(ctx context.Context, userID int64) (*model.User, error) {
						return &model.User{
							UserID:    1,
							IsPremium: false,
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					GetRelatedUserCacheLenFunc: func(ctx context.Context, userID int64) (int64, error) {
						return 6, errors.New("err")
					},
				},
			},
			args: args{
				req: model.SwipeRequest{
					UserID:       1,
					SwipedUserID: 2,
					SwipeStatus:  -1,
				},
			},
			wantErr: true,
		},
		{
			name: "case error cache set",
			fields: fields{
				repoDB: &db.RepoMock{
					GetUserByIDFunc: func(ctx context.Context, userID int64) (*model.User, error) {
						return &model.User{
							UserID:    1,
							IsPremium: false,
						}, nil
					},
				},
				repoCache: &cache.RepoMock{
					GetRelatedUserCacheLenFunc: func(ctx context.Context, userID int64) (int64, error) {
						return 6, nil
					},
					SetRelatedUserCacheFunc: func(ctx context.Context, userID int64, data model.UserRelation) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				req: model.SwipeRequest{
					UserID:       1,
					SwipedUserID: 2,
					SwipeStatus:  -1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				RepoDB:    tt.fields.repoDB,
				RepoCache: tt.fields.repoCache,
			}
			gotErr := u.Swipe(context.Background(), tt.args.req)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Swipe() error = %v, wantErr = %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}
