package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/hassaanaliw/go-web-notes/model"
)

// Registers all the various JSON API endpoints with the main app router
// created in main.go
func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/notes/all/", AllNotesJSONRoute).Methods("GET")
	router.HandleFunc("/api/v1/notes/sample/", SampleJSONRoute).Methods("GET")
}

// URL: http://localhost:PORT/api/v1/notes/all/
// Queries the database for all the notes and returns a JSON response
// containing the following details:
//{
//	title: "Test Note",
//	description: "",
//	text: "This is a Test",
//	note_type: "Normal",
//	date_published: "2018-06-11T17:34:23.131841999+05:00",
//	date_edited: "0001-01-01T00:00:00Z",
//	categories: [ ],
//	is_edited: false
//}
func AllNotesJSONRoute(w http.ResponseWriter, r *http.Request) {
	f := map[string]interface{}{
		"Notes": []model.Note{
		},
	}
	ServeJson(&w, r, f)

}

// URL: http://localhost:PORT/api/v1/notes/sample/
// Returns the JSON of a sample Note, used for testing purposes
func SampleJSONRoute(w http.ResponseWriter, r *http.Request) {
	note := model.CreateNote("Test Note", "This is a Test", "Normal", []model.Category{})
	ServeJson(&w, r, note)

}

// Sets the appropriate MIME types for a JSON response and modifies the headers
// Encodes the request string as a JSON object and writes it to the response
func ServeJson(w *http.ResponseWriter, r *http.Request, jsonResponse interface{}) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(*w).Encode(jsonResponse)
	if err != nil {
		// Log error if any but don't kill thr request
		fmt.Println(err)
	}
}
