package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WebClientModule handles the registration of the webclient module
func WebClientModule(r *mux.Router) {
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))
}
