package models

import "time"

// LoginModul ...
type LoginModul struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse ...
type TokenResponse struct {
	Token string `json:"token"`
}

// CreateUserModul ...
type CreateUserModul struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	User_type string `json:"user_type" binding:"required"`
}

// User ...
type User struct {
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	User_type string     `json:"user_type" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
