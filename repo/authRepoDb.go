package repo

import (
	"learn-go-goroutine/models"

	"gorm.io/gorm"
)

type authRepoDb struct {
	db *gorm.DB
}

func NewAuthRepoDb(db *gorm.DB) *authRepoDb{
	return &authRepoDb{db:db}
}

func (a *authRepoDb) GetUserByEmail(email string) (*models.User, error){
	user := models.User{Email: email}

	result := a.db.First(&user)

	if result.Error != nil{
		return nil,result.Error
	}

	return &user,nil
}
