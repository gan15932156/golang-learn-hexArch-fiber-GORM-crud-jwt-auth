package repo

import "learn-go-goroutine/models"

type AuthRepo interface {
	GetUserByEmail(string) (*models.User, error)
}