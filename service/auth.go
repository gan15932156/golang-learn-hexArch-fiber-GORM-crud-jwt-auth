package service

import "learn-go-goroutine/types"

type AuthService interface {
	SignIn(*types.SignInRequest) (*types.AuthResponse,error)
}