package link

import (
	"fmt"
	"go-serv/demo/pkg/request"
	"go-serv/demo/pkg/responses"
	"net/http"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)

		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		responses.Json(w, createdLink, 201)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) FindByUrl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}
		url, err := handler.LinkRepository.GetByUrl(body.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		responses.Json(w, url, http.StatusOK)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//func
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	link := LinkHandler{LinkRepository: deps.LinkRepository}

	router.HandleFunc("POST /create", link.Create())
	router.HandleFunc("PATCH /link/{id}", link.Update())
	router.HandleFunc("DELETE /link/{id}", link.Delete())
	router.HandleFunc("GET /{hash}", link.GoTo())
	router.HandleFunc("GET /link/find/url", link.FindByUrl())
}
