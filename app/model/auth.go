package model

type (
	Auth struct {
		User User `json:"user"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RegisterRequest struct {
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
	}
	// LogoutRequest struct {
	// 	Token string `json:"token" validate:"required"`
	// }
)
