/*
Defines the models used in this web application as structs.
Hassaan Ali Wattoo <hawattoo@umich.edu>
*/

package model

import "time"

type Note struct {
	Title         string
	Description   string
	Text          string
	NoteType      string
	DatePublished time.Time
	DateEdited    time.Time
	Categories    []Category
}

type Category struct {
	Name string
}
