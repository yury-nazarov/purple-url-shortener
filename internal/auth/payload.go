package auth

// Декодируем пользовательский JSON
type LoginRequest struct {
	// В валидаторе email - это инструкция для парсинга Email string (https://github.com/go-playground/validator)
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}


type RegisterResponse struct {
	Token string `json:"token"`
}