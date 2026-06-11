package response

type RegisterResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MidName string `json:"mid_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
}

type RefreshTokenResponse struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	MidName string `json:"mid_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
