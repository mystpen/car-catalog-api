package errorres

import (
	"fmt"
	"net/http"

	"github.com/mystpen/car-catalog-api/pkg/jsonutil"
	"github.com/mystpen/car-catalog-api/pkg/logger"
)

func LogError(r *http.Request, err error) {
	logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := jsonutil.Envelope{"error": message}

	err := jsonutil.WriteJSON(w, status, env, nil)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}
