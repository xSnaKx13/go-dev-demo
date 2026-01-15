package main

import (
	"fmt"
	"go-serv/demo/configs"
	"go-serv/demo/internal/auth"
	"go-serv/demo/internal/link"
	"go-serv/demo/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(&conf)
	router := http.NewServeMux()

	newLinkRepository := link.NewLinkRepository(db)

	//handlers
	link.NewLinkHandler(router, &link.LinkHandlerDeps{LinkRepository: newLinkRepository})
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{Config: &conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server listening on port 8081...")
	server.ListenAndServe()
}
