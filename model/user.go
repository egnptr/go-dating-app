package model

type User struct {
	UserID    int64  `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Email     string `json:"email,omitempty"`
	IsPremium bool   `json:"is_premium,omitempty"`
}

type UserRelation struct {
	UserID      int64 `json:"id"`
	SwipeStatus int   `json:"swipe_status"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SubscribeRequest struct {
	UserID    int64 `json:"user_id"`
	Subscribe bool
}

type SwipeRequest struct {
	UserID       int64 `json:"user_id"`
	SwipedUserID int64 `json:"swiped_user_id"`
	SwipeStatus  int   `json:"swipe_status"`
}

type GetRelatedUserRequest struct {
	UserID int64 `json:"user_id"`
}
