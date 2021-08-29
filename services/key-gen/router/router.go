package router

import (
	"github.com/gorilla/mux"
	"github.com/robatipoor/short-link/services/key-gen/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping).Methods("GET")
	r.HandleFunc("/getkey", handler.GetKey).Methods("GET")
	r.HandleFunc("/usekey/{key}", handler.UseKey).Methods("GET")
	r.HandleFunc("/{key}", handler.DelKey).Methods("DELETE")
	return r
}
