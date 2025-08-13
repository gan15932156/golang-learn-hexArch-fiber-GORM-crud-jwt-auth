package service

import (
	"errors"
	customvalidate "learn-go-goroutine/customValidate"
	"learn-go-goroutine/repo"
	"learn-go-goroutine/types"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repo.UserRepo
	validator *validator.Validate
}

func NewUserService(repo repo.UserRepo,validator *validator.Validate) *userService{
	return &userService{repo:repo,validator:validator}
}

func (u *userService) Register(user *types.User) (*types.RegisterResponse,error){
	errrorMessages := customvalidate.Validate(u.validator,user)

	if len(errrorMessages) > 0 {
		errMsg := strings.Join(errrorMessages, ", ")
		return nil,errors.New(errMsg)
	}

	hash, hashedError := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if hashedError != nil {
        return nil,hashedError
    }

	user.Password = string(hash)

	_,createError := u.repo.Create(*user)

	if createError != nil {
        return nil,createError
    }


	return nil,nil
}