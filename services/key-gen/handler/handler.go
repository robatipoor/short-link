package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robatipoor/short-link/services/key-gen/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}
func GetKey(w http.ResponseWriter, r *http.Request) {
	exp := r.FormValue("exp")
	key, err := service.GetKey(exp)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed get key", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(key))

}

func UseKey(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	key := m["key"]
	exp := r.FormValue("exp")
	err := service.UseKey(key, exp)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed use key", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func DelKey(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	key := m["key"]
	log.Println("delete key : ", key)
	err := service.DelKey(key)
	if err != nil {
		log.Println("delete key service failed details : ",err)
		http.Error(w, "failed delete key", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
