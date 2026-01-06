package main

import (
	"fmt"
	"go-serv/demo/configs"
	"go-serv/demo/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server listening on port 8081...")
	server.ListenAndServe()
}
