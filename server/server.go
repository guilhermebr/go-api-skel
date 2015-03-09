package server

import (
	"github.com/guilhermebr/go-api-skel/modules/task"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers() {
	fmt.Println("Register handlers")
	r := mux.NewRouter()

	task.RegisterRoute(r)
	http.Handle("/task/", r)
}
