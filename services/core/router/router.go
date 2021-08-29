package router

import (
	"github.com/gorilla/mux"
	"github.com/robatipoor/short-link/services/core/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping)
	r.HandleFunc("/{key}", handler.Redirect).Methods("GET")
	r.HandleFunc("/", handler.CreateNewLink).Methods("POST")
	r.HandleFunc("/{key}", handler.DeleteLink).Methods("DELETE")
	return r
}
