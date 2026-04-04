package auth

import (
	"net/http"

	"adv-demo/configs"
	"adv-demo/pkg/req"
)

// Структура используемачя для передачи зависимости в компонент
// Набор полей у AuthHandlerDeps и AuthHandler
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
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[RegisterRequest](&w, r)
		if err == nil {
			return
		}
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TBD: Добавить валидацию
		_, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
	}
}
