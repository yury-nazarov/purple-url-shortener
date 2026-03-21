package main

import (
	//"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/hello"

	"fmt"
	"net/http"
)

func main() {
	//conf := configs.LoadConfig()

	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	// Сделать Auth модуль
	// Добавить два хендлера /auth/register, /auth/login
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listining on port 8081")
	server.ListenAndServe()

}
