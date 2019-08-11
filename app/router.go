package app

import (
	"net/http"

	"github.com/george-kj/krsc/handler"
	"github.com/gorilla/mux"
)

//Routes creates the handlers, routes
func Routes() *mux.Router {
	// Status routes.
	r := mux.NewRouter()
	r.Handle("/status", handler.ResponseHandler(handler.Status)).Methods("Get")

	r.Use(Middleware)
	return r
}

// Middleware handles authentication for all routes.
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

	})
}
