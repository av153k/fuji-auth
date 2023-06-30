package models

// Register struct to describe the Regsiter model
type Register struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	UserRole string `json:"user_role"`
}

// Login struct to describe the Login model
type Login struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
