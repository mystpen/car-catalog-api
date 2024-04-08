package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mystpen/car-catalog-api/pkg/errorres"
)

func (h *Handler) Routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(errorres.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(errorres.MethodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/cars", h.listCarsHandler)
	router.HandlerFunc(http.MethodPost, "/cars", h.addCarInfoHandler)
	router.HandlerFunc(http.MethodPatch, "/cars/:id", h.updateCarInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/cars/:id", h.deleteCarInfoHandler)


	return router
}
