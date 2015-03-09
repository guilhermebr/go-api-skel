package task

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/task/", errorHandler(ListTasks)).Methods("GET")
	r.HandleFunc("/task/", errorHandler(CreateTask)).Methods("POST")
	r.HandleFunc("/task/"+"{id}", errorHandler(GetTask)).Methods("GET")
	r.HandleFunc("/task/"+"{id}", errorHandler(UpdateTask)).Methods("PUT")

}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// errorHandler wraps a function returning an error by handling the error and returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as described for every type.
// If the error is of another type, it is considered as an internal error and its message is logged.
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}
