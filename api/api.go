package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
)

// Registers all the various JSON API endpoints with the main app router
// created in main.go
func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/notes/all/", AllNotesJSONRoute).Methods("GET")
}

// http://localhost:PORT/api/v1/notes/all
func AllNotesJSONRoute(w http.ResponseWriter, r *http.Request) {
	response := "{'notes':[]}"§
	ServeJson(&w, r, &response)

}

// Sets the appropriate MIME types for a JSON response and modifies the headers
// Encodes the request string as a JSON object and writes it to the response
func ServeJson(w *http.ResponseWriter, r *http.Request, jsonResponse *string) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(*w).Encode(*jsonResponse)

	if err != nil {
		// Log error if any but don't kill thr request
		fmt.Println(err)
	}
}
