package repo

import (
	"learn-go-goroutine/models"
	"learn-go-goroutine/types"

	"gorm.io/gorm"
)

type userRepoMockDb struct {
}

var users = []models.User{}

func NewUserRepoMockDb() *userRepoMockDb {
	return &userRepoMockDb{}
}

func (u *userRepoMockDb) Create(user types.User) (*models.User,error) {
	newUser := models.User{Model: gorm.Model{ID: 1},Name: user.Name,Email: user.Email,Password: user.Password}
	users = append(users,newUser)
	return &newUser,nil
}
func (u *userRepoMockDb) GetAll() (*[]models.User,error){
	return &users,nil
}