package repo

import (
	"learn-go-goroutine/models"
	"learn-go-goroutine/types"
)

type UserRepo interface {
	Create(types.User) (*models.User,error)
	GetAll() (*[]models.User,error)
}