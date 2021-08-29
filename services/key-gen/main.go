package main

import (
	"net/http"

	"github.com/robatipoor/short-link/services/key-gen/config"
	"github.com/robatipoor/short-link/services/key-gen/router"
)

func main() {
	defer config.SessionDB.Close()
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+config.ServerPort, r)
}
