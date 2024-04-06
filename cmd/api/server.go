package main

import "github.com/mystpen/car-catalog-api/internal/delivery/http"


type httpserver struct{
	handler *http.Handler
}

func (s httpserver) New(handler *http.Handler) httpserver{
	return httpserver{
		handler: handler,
	}
}