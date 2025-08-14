package service

import (
	"errors"
	customvalidate "learn-go-goroutine/customValidate"
	"learn-go-goroutine/repo"
	"learn-go-goroutine/types"
	"learn-go-goroutine/utils"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
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

	hash, hashedError := utils.HashPassword(user.Password)

	if hashedError != nil {
        return nil,hashedError
    }

	user.Password = string(hash)

	insertedUser,createError := u.repo.Create(*user)

	if createError != nil {
        return nil,createError
    }

	jwtPayload := utils.JwtPayload{
		Sub: insertedUser.Name,
		Iss: "App",
		Exp: time.Now().Add(time.Hour).Unix(),
		Iat: time.Now().Unix(),
	}

	token,errorToken := utils.SignJwtToken(&jwtPayload)

	if errorToken != nil {
        return nil, errorToken
    }

	return &types.RegisterResponse{
		User: types.ResponseUser{
			Id: insertedUser.ID,
			Name: insertedUser.Name,
			Email: insertedUser.Email},
		Auth: types.AuthResponse{Token: token}},nil
}

func (u *userService) GetUsers() (*[]types.ResponseUser,error){
	users := []types.ResponseUser{}

	resultUsers,err := u.repo.GetAll();

	if  err != nil{
		return &[]types.ResponseUser{},err
	}

	for _, user := range *resultUsers {
		users = append(users,types.ResponseUser{Id: user.ID,Name: user.Name,Email: user.Email})
		
	}

	return &users,nil
}

func (u *userService) UpdateUser(user types.UpdateUser,id uint) error{
	errrorMessages := customvalidate.Validate(u.validator,user)

	if len(errrorMessages) > 0 {
		errMsg := strings.Join(errrorMessages, ", ")
		return errors.New(errMsg)
	}

	err := u.repo.Update(user,id)

	if err != nil{
		return err
	}

	return nil
}