package service

import (
	"errors"
	"fmt"
	customvalidate "learn-go-goroutine/customValidate"
	"learn-go-goroutine/repo"
	"learn-go-goroutine/types"
	"strings"

	"github.com/go-playground/validator/v10"
)

type authService struct {
	authRepo repo.AuthRepo
	validator *validator.Validate
}

func NewAuthService (authRepo repo.AuthRepo,validator *validator.Validate) *authService{
	return &authService{authRepo: authRepo,validator: validator}
}

func (a *authService) SignIn(cred *types.SignInRequest) (*types.AuthResponse,error){
	errrorMessages := customvalidate.Validate(a.validator,cred)

	if len(errrorMessages) > 0 {
		errMsg := strings.Join(errrorMessages, ", ")
		return nil,errors.New(errMsg)
	}

	user,err := a.authRepo.GetUserByEmail(cred.Email)

	if err != nil{
		return nil,err
	}

	fmt.Println(user)
	
	return nil,nil
}