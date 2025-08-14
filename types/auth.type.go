package types

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=2,max=8"`
}

type AuthResponse struct {
	Token string `json:"token"`
}