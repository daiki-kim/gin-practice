package dto

type SignUpUserInput struct {
	Email    string `json:"email" binding:"required,email"` // email: email形式であるかのbinding
	Password string `json:"password" binding:"required,min=8"`
}

type LogInUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
