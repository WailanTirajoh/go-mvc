package model

type (
	Auth struct {
		User User `json:"user"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LogoutRequest struct {
		Token string `json:"token" validate:"required"`
	}
)
