package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinsudut/single-fizz-buzz/app/usecase"
)

type handler struct {
	usecase usecase.UsecaseItf
}

func Init() *handler {
	return &handler{
		usecase: usecase.Init(),
	}
}

func (h handler) RegisterHandlers(router *mux.Router) *mux.Router {
	// Define a HTTP endpoint called GET /range-fizzbuzz
	router.HandleFunc("/range-fizzbuzz", h.handleSingleFizzBuzzWithRange).Methods(http.MethodGet)

	return router
}
