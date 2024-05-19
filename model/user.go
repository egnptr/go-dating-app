package model

type User struct {
	UserID    int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	IsPremium bool   `json:"is_premium"`
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
