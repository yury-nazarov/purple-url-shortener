package main

import (
	"adv-demo/internal/hello"

	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listining on port 8081")
	server.ListenAndServe()

}
