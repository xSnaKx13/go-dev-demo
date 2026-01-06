package auth

import (
	"go-serv/demo/configs"
	"go-serv/demo/pkg/request"
	"go-serv/demo/pkg/responses"
	"log"
	"net/http"
)

type AuthHandler struct {
	*AuthHandlerDeps
}

type AuthHandlerDeps struct {
	*configs.Config
}

func (handler *AuthHandler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequestPayload](w, r)
		if err != nil {
			return
		}
		log.Printf("%+v\n", body)
	}
}

func (handler *AuthHandler) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequestPayload](w, r)
		if err != nil {
			return
		}
		log.Printf("%+v\n", body)
		responses.Json(w, 201, "Register successful!")
	}
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		AuthHandlerDeps: deps,
	}
	router.HandleFunc("POST /login", handler.LoginHandler())
	router.HandleFunc("POST /register", handler.RegisterHandler())
}
