package api

import (
	"time"
)

type createUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type updateUserParams struct {
	ID        int64     `json:"id" validate:"required"`
	Password  string    `json:"password" validate:"required,gte=8"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

type authParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type authResponse struct {
	User  *userResponse `json:"user"`
	Token string        `json:"token"`
}

type genericResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
