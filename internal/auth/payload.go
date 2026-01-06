package auth

type LoginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type LoginResponsePayload struct {
	Token string `json:"token"`
}

type RegisterRequestPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type RegisterResponsePayload struct {
	Token string `json:"token"`
}
