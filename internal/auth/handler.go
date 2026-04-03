package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"adv-demo/configs"
	"adv-demo/pkg/res"
)

// Структура используемачя для передачи зависимости в компонент
// Набор полей у AuthHandlerDeps и AuthHandler
// TBD: Выглядит как будто эта структурка может быть приватной
type AuthHandlerDeps struct {
	*configs.Config
}

// Структура используемая для функции конструктора
// Набор полей у AuthHandlerDeps и AuthHandler
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		// Передаем весь конфиг в учебных целях для упрощения.
		// Иначе было бы достаточно только токена
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/regiser", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

// Разберем подробно, так как быбло не очевидно.
// Функция Register() возвращает функцию типа HandlerFunc - а точнее замыкание, которое реализует ее интерфейс.
// В момент возврата, анонимная функция не выполняется
// Она будет выполнена, когда дернут за ручку
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("registeration new user")
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Прочитать Body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}
		if payload.Email == "" {
			res.Json(w, "Email required", 402)
			return
		}
		if payload.Password == "" {
			res.Json(w, "Password required", 402)
			return
		}
		fmt.Println(payload.Email)
		data := LoginResponse{
			Token: handler.Config.Auth.Secret,
		}
		res.Json(w, data, 200)
	}
}
