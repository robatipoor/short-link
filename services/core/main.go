package main

import (
	"net/http"

	"github.com/robatipoor/short-link/services/core/config"
	"github.com/robatipoor/short-link/services/core/router"
)

func init() {
}

func main() {
	defer config.SessionDB.Close()
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+config.ServerPort, r)
}
