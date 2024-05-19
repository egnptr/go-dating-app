package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/egnptr/dating-app/model"
	"github.com/egnptr/dating-app/usecase"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	type fields struct {
		service usecase.Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "case success",
			fields: fields{
				service: &usecase.UsecasesMock{
					CreateUserFunc: func(ctx context.Context, req model.User) error {
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"username": "abc",
						"password": "test123",
						"full_name": "test",
						"email": "test1"
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 200,
		},
		{
			name: "case error",
			fields: fields{
				service: &usecase.UsecasesMock{
					CreateUserFunc: func(ctx context.Context, req model.User) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"username": "abc",
						"password": "test123",
						"full_name": "test",
						"email": "test1"
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Usecase: tt.fields.service,
			}
			c.SignUp(tt.args.w, tt.args.r)
			if recorder, ok := tt.args.w.(*httptest.ResponseRecorder); ok && recorder != nil {
				assert.Equal(t, tt.wantCode, recorder.Code)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	type fields struct {
		service usecase.Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "case success",
			fields: fields{
				service: &usecase.UsecasesMock{
					LoginFunc: func(ctx context.Context, req model.LoginRequest) error {
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"username": "abc",
						"password": "test123"
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 200,
		},
		{
			name: "case error",
			fields: fields{
				service: &usecase.UsecasesMock{
					LoginFunc: func(ctx context.Context, req model.LoginRequest) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"username": "abc",
						"password": "test123"
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Usecase: tt.fields.service,
			}
			c.LoginUser(tt.args.w, tt.args.r)
			if recorder, ok := tt.args.w.(*httptest.ResponseRecorder); ok && recorder != nil {
				assert.Equal(t, tt.wantCode, recorder.Code)
			}
		})
	}
}

func TestUpdateSubscription(t *testing.T) {
	type fields struct {
		service usecase.Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "case success",
			fields: fields{
				service: &usecase.UsecasesMock{
					UpdateSubscriptionFunc: func(ctx context.Context, req model.SubscribeRequest) error {
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/unsubscribe", strings.NewReader(`{
						"user_id": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 200,
		},
		{
			name: "case error",
			fields: fields{
				service: &usecase.UsecasesMock{
					UpdateSubscriptionFunc: func(ctx context.Context, req model.SubscribeRequest) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/subscribe", strings.NewReader(`{
						"user_id": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Usecase: tt.fields.service,
			}
			c.UpdateSubscription(tt.args.w, tt.args.r)
			if recorder, ok := tt.args.w.(*httptest.ResponseRecorder); ok && recorder != nil {
				assert.Equal(t, tt.wantCode, recorder.Code)
			}
		})
	}
}

func TestGetProfiles(t *testing.T) {
	type fields struct {
		service usecase.Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "case success",
			fields: fields{
				service: &usecase.UsecasesMock{
					GetProfilesFunc: func(ctx context.Context, req model.GetRelatedUserRequest) ([]model.User, error) {
						return []model.User{}, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"user_id": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 200,
		},
		{
			name: "case error",
			fields: fields{
				service: &usecase.UsecasesMock{
					GetProfilesFunc: func(ctx context.Context, req model.GetRelatedUserRequest) ([]model.User, error) {
						return []model.User{}, errors.New("err")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{
						"user_id": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Usecase: tt.fields.service,
			}
			c.GetProfiles(tt.args.w, tt.args.r)
			if recorder, ok := tt.args.w.(*httptest.ResponseRecorder); ok && recorder != nil {
				assert.Equal(t, tt.wantCode, recorder.Code)
			}
		})
	}
}

func TestSwipe(t *testing.T) {
	type fields struct {
		service usecase.Usecases
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "case success",
			fields: fields{
				service: &usecase.UsecasesMock{
					SwipeFunc: func(ctx context.Context, req model.SwipeRequest) error {
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"user_id": 1,
						"swiped_user_id": 2,
						"swipe_status": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 200,
		},
		{
			name: "case error",
			fields: fields{
				service: &usecase.UsecasesMock{
					SwipeFunc: func(ctx context.Context, req model.SwipeRequest) error {
						return errors.New("err")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{
						"user_id": 1,
						"swiped_user_id": 2,
						"swipe_status": 1
					}`))
					request.Header.Set("Content-Type", "application/json")
					return request
				}(),
			},
			wantCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				Usecase: tt.fields.service,
			}
			c.Swipe(tt.args.w, tt.args.r)
			if recorder, ok := tt.args.w.(*httptest.ResponseRecorder); ok && recorder != nil {
				assert.Equal(t, tt.wantCode, recorder.Code)
			}
		})
	}
}
