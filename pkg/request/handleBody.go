package request

import (
	"go-serv/demo/pkg/responses"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	body, err := Decode[T](r)
	if err != nil {
		responses.Json(w, http.StatusBadRequest, "Invalid request payload")
		return body, err
	}
	err = IsValidate(body)
	if err != nil {
		responses.Json(w, http.StatusBadRequest, "Invalid request payload")
		return body, err
	}
	return body, nil
}
