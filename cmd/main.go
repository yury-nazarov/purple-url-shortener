package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/pkg/db"

	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.New(conf)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listining on port 8081")
	server.ListenAndServe()

}
