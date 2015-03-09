package main

import (
	"net/http"

	"github.com/guilhermebr/go-api-skel/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
