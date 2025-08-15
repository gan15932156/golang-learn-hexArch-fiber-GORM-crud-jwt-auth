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
	"golang.org/x/crypto/bcrypt"
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

	user,errGetUser := a.authRepo.GetUserByEmail(cred.Email)

	if errGetUser != nil{
		return nil,errGetUser
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(cred.Password))

	if errCompare != nil{
		return nil,errors.New("invalid credential")
	}

	jwtPayload := utils.JwtPayload{
		Sub: user.Name,
		Iss: "App",
		Exp: time.Now().Add(time.Hour).Unix(),
		Iat: time.Now().Unix(),
	}

	token,errorToken := utils.SignJwtToken(&jwtPayload)

	if errorToken != nil {
        return nil, errorToken
    }
	
	return &types.AuthResponse{Token: token},nil
}