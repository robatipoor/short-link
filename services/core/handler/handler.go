package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robatipoor/short-link/services/core/domain/request"
	"github.com/robatipoor/short-link/services/core/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	var req request.Redirect
	m := mux.Vars(r)
	key := m["key"]
	req.ShortUrl = key
	resp, err := service.Redirect(&req)
	if err != nil {
		http.Error(w, "redirect failed", http.StatusInternalServerError)
		return
	}
	log.Printf("redirect %s to %s \n", key, resp.Link)
	http.Redirect(w, r, resp.Link, http.StatusSeeOther)
}

func CreateNewLink(w http.ResponseWriter, r *http.Request) {
	var req request.CreateNewUrl
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("decode request failed error details : ", err.Error())
		http.Error(w, "create new link decode request failed ", http.StatusInternalServerError)
		return
	}
	resp, err := service.CreateNewLink(&req)
	if err != nil {
		log.Println("create new link failed error details : ", err.Error())
		http.Error(w, "create new link failed ", http.StatusInternalServerError)
		return
	}
	log.Println("success create new link ", resp)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Link))
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	m := mux.Vars(r)
	key := m["key"]
	err := service.DeleteLink(key)
	if err != nil {
		http.Error(w, "delete link failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
