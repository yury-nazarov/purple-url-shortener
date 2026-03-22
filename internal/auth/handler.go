package auth

import (
	"fmt"
	"net/http"
)

type ConfigProvider interface {
	GetAuthSecret() string
}

// Структура используемачя для передачи зависимости в компонент
// Набор полей у AuthHandlerDeps и AuthHandler может отличаться.
type AuthHandlerDeps struct {
	Config ConfigProvider
}

// Структура используемая для функции конструктора
// Набор полей у AuthHandlerDeps и AuthHandler может отличаться.
type authHandler struct {
	Config ConfigProvider
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		// Объект Config должнен соответствовать интерфейсу
		// ConfigProvider (должны быть описаные методы)
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/regiser", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("registeration new user")
	}
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		secret := handler.Config.GetAuthSecret()
		fmt.Println("login with secret:", secret)
	}
}
