package auth

import (
	"fmt"
	"net/http"

	"adv-demo/configs"
	self_hash "adv-demo/pkg/hash"
	"adv-demo/pkg/mailer"
	"adv-demo/pkg/req"
	"adv-demo/pkg/res"
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
	store map[string]string
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		// Передаем весь конфиг в учебных целях для упрощения.
		// Иначе было бы достаточно только токена
		Config: deps.Config,
		store:  make(map[string]string),
	}

	router.HandleFunc("POST /auth/regiser", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /send", handler.SendEmail())
	router.HandleFunc("GET /verify/{hash}", handler.VerifyHash())
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

func (handler *AuthHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Принимает email и валидирует его
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}
		// Готовим хеш
		userHash, err := self_hash.Gen(body.Email)
		if err != nil {
			res.Json(w, "Hash generate Internal error", 500)
			return
		}

		// Формируем информацию о пользователе
		// В будущем будем брать ее из БД
		user := mailer.UserInfo{
			Name:         "Yury",
			Email:        body.Email,
			EmailPasword: "31pHOmYwf1HhV7o7kcI7",
			SmtpServer:   "smtp.mail.ru",
			SmtpPort:     "587",
			Hash:         userHash,
		}
		// Сохраняем пару для проверки
		handler.store[userHash] = body.Email

		// Отправляем ссылку о подтверждении регистрации пользователю
		err = mailer.Send(user)
		if err != nil {
			res.Json(w, "ERROR: Mailer has problem", 500)
			return
		}
	}
}

func (handler *AuthHandler) VerifyHash() http.HandlerFunc {
	// TODO: В mail.ru вместо нормального хеша приходит какая то херня
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		value, ok := handler.store[hash]
		if ok {
			fmt.Printf("INFO: email %s was conferm:", value)
		} else {
			fmt.Println("ERROR: email not found")
		}

	}
}
