/*

Main entry point for our web app. Initialises the router and the mux ware,
sets the config variables, and serves the web application

Hassaan Ali Wattoo <hawattoo@umich.edu>

*/

package main

import (
	"fmt"
	"github.com/bndr/gotabulate"
)

// Main entry point for this application
func main() {
	loadStatus, config := loadConfig()

	if !loadStatus {
		// Kill if we can't load config file
		fmt.Print("Error Loading Configuration File")
		return
	}

	configOutputHeaders := []string{"Config Variables", "Config Value"}
	row_debug := []interface{}{"Debug", config.Debug}
	row_db_url := []interface{}{"Database URL", config.DatabaseURL}
	row_port := []interface{}{"Port", config.Port}

	if config.Debug {
		tableOutput(configOutputHeaders, [][]interface{}{row_debug, row_db_url, row_port})
	}

}

// Helper function to output data in a nice tabular format
// headers is a slice of table headers, rows are the rows to print
func tableOutput(headers []string, rows [][]interface{}) {

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
