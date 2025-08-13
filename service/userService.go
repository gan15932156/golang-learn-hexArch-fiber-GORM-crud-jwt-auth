package service

import (
	"errors"
	"learn-go-goroutine/config"
	customvalidate "learn-go-goroutine/customValidate"
	"learn-go-goroutine/repo"
	"learn-go-goroutine/types"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

	insertedUser,createError := u.repo.Create(*user)

	if createError != nil {
        return nil,createError
    }

	secret := config.Config("SECRET")

	if len(secret) == 0{
		return nil,errors.New("Error")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": insertedUser.Name,                  
		"iss": "kuay-app",                 
		"exp": time.Now().Add(time.Hour).Unix(), 
		"iat": time.Now().Unix(),                
	})

	tokenString, signedError := claims.SignedString([]byte(secret))

	if signedError != nil {
        return nil, signedError
    }

	return &types.RegisterResponse{
		User: types.ResponseUser{
			Id: insertedUser.ID,
			Name: insertedUser.Name,
			Email: insertedUser.Email},
		Auth: types.AuthResponse{Token: tokenString}},nil
}