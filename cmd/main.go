package main

import (
	"fmt"
	"go-serv/demo/configs"
	"go-serv/demo/internal/auth"
	"go-serv/demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(&conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{Config: &conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server listening on port 8081...")
	server.ListenAndServe()
}
