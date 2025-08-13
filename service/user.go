package service

import "learn-go-goroutine/types"

type UserService interface {
	Register(*types.User) (*types.RegisterResponse,error)
	UpdateUser(types.UpdateUser,uint) error
	GetUsers() (*[]types.ResponseUser,error)
}