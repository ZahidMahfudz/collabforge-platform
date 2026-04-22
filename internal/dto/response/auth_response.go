package response

type UserDataResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	Token string           `json:"token"`
	User  UserDataResponse `json:"user"`
}
