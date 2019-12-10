package handlers

import (
	"chatterbox/server/compiled"
	"net/http"

	"github.com/gorilla/mux"
)

// PrecompiledWebClientModule serves the precompiled web client
func PrecompiledWebClientModule(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, _ := compiled.IndexHTML()
		w.Write(html)
	})
}
