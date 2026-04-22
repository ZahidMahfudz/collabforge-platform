package request

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type GoogleLoginRequest struct {
	IDToken string `json:"id_token" validate:"required"`
}
