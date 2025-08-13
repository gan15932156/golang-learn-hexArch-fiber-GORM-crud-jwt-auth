package repo

import (
	"errors"
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

func (u *userRepoDb) Create(user types.User) (models.User,error){
	newUser := models.User{Name: user.Name,Email: user.Email,Password: user.Password}
	result := u.db.Create(&newUser)
	if result.Error != nil{
		return models.User{},result.Error
	}
	return newUser,nil
}

func (u *userRepoDb) GetAll() (*[]models.User,error){
	users := []models.User{}
	result := u.db.Find(&users)
	if result.Error != nil{
		return nil,result.Error
	}
	return &users,nil
}

func (u *userRepoDb) Update(user types.UpdateUser,id uint) error{
	updatedUser := models.User{Model: gorm.Model{ID: id}}
	
	if result := u.db.First(&updatedUser); result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}

	updatedUser.Name = user.Name

	if result := u.db.Save(&updatedUser); result.Error != nil{
		return result.Error
	}

	return nil
	
}