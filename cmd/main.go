package main

import (
	"adv-demo/configs"
	"adv-demo/internal/auth"
	"adv-demo/internal/link"
	"adv-demo/pkg/db"

	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.New(conf)
	router := http.NewServeMux()

	// Repository
	linlRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linlRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listining on port 8081")
	server.ListenAndServe()

}
