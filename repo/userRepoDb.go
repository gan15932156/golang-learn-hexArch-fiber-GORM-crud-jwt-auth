package repo

import (
	"learn-go-goroutine/models"
	"learn-go-goroutine/types"

	"gorm.io/gorm"
)

type userRepoDb struct {
	db *gorm.DB
}

func NewUserRepoDb(db *gorm.DB) *userRepoDb{
	return &userRepoDb{db:db}
}

func (u *userRepoDb) Create(user types.User) error{
	result := u.db.Create(&user)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func (u *userRepoDb) GetAll() (*[]models.User,error){
	users := []models.User{}
	result := u.db.Find(&users)
	if result.Error != nil{
		return nil,result.Error
	}
	return &users,nil
}