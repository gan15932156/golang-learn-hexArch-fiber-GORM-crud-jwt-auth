package utils

import (
	"errors"
	"learn-go-goroutine/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte,error){
	hash, hashedError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash,hashedError
}

type JwtPayload struct{
	Sub string
	Iss string
	Exp int64
	Iat int64
}
              
func SignJwtToken(payload *JwtPayload) (string,error) {
	secret := config.Config("SECRET")

	if len(secret) == 0{
		return "",errors.New("Error")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.Sub,                  
		"iss": payload.Iss,                 
		"exp": payload.Exp, 
		"iat": payload.Iat,                
	})

	tokenString, signedError := claims.SignedString([]byte(secret))

	if signedError != nil {
        return "", signedError
    }

	return tokenString,nil
}