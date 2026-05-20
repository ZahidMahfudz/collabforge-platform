package request

type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=50"`
	MidName   string `json:"mid_name" validate:"omitempty,max=50"`
	LastName  string `json:"last_name" validate:"required,min=3,max=50"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,password"`
}