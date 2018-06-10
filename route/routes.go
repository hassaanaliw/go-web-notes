package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// This handler serves CSS stylesheets from the static directory
// It's main purpose is to set the StyleSheet MIME-Type since browsers
// are strict about enforcing them
func CSSHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url := fmt.Sprintf("static/css/%s", vars["filename"])

	w.Header().Set("Content-Type", "text/css")

	http.ServeFile(w, r, url)
}

// This handler serves CSS stylesheets from the static directory
// It's main purpose is to set the StyleSheet MIME-Type since browsers
// are strict about enforcing them
func JSHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	url := fmt.Sprintf("static/js/%s", vars["filename"])

	w.Header().Set("Content-Type", "text/javascript")

	http.ServeFile(w, r, url)
}
