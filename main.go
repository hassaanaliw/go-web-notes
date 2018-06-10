/*

Main entry point for our web app. Initialises the router and the mux ware,
sets the config variables, and serves the web application

Hassaan Ali Wattoo <hawattoo@umich.edu>

*/

package main

import (
	"fmt"
	"github.com/bndr/gotabulate"
	"github.com/hassaanaliw/go-web-notes/config"
	"github.com/gorilla/mux"
	"github.com/hassaanaliw/go-web-notes/route"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

// Main entry point for this application
func main() {

	// Load Configuration
	loadStatus, configuration := config.LoadConfig()

	if !loadStatus {
		// Kill if we can't load config file
		fmt.Print("Error Loading Configuration File")
		return
	}

	// Initialise the configuration logger
	configOutputHeaders := []string{"Config Variables", "Config Value"}
	row_debug := []interface{}{"Debug", configuration.Debug}
	row_db_url := []interface{}{"Database URL", configuration.DatabaseURL}
	row_port := []interface{}{"Port", configuration.Port}

	if configuration.Debug {
		TableOutput(configOutputHeaders, [][]interface{}{row_debug, row_db_url, row_port})
	}

	// Setup Router
	router := mux.NewRouter()
	router.HandleFunc("/", route.Index)
	router.HandleFunc("/static/css/{filename}", route.CSSHandler)
	router.HandleFunc("/static/js/{filename}", route.JSHandler)

	// Serve the HTTP Application
	fmt.Printf("Serving web app on url: http://localhost:%d\n", configuration.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port),
		handlers.LoggingHandler(os.Stdout, router))

	if err != nil {
		fmt.Println(err)
	}

}

// Helper function to output data in a nice tabular format
// headers is a slice of table headers, rows are the rows to print
func TableOutput(headers []string, rows [][]interface{}) {

	// Create an object from 2D interface array
	t := gotabulate.Create(rows)

	// Set the Headers (optional)
	t.SetHeaders(headers)

	// Set the Empty String (optional)
	t.SetEmptyString("None")

	// Set Align (Optional)
	t.SetAlign("right")

	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))

}
