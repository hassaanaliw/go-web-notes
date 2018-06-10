/*
Defines the models used in this web application as structs.
Hassaan Ali Wattoo <hawattoo@umich.edu>
*/

package model

import (
	"time"
	"fmt"
)

type Note struct {
	Title         string
	Description   string
	Text          string
	NoteType      string
	DatePublished time.Time
	DateEdited    time.Time
	Categories    []Category
	IsEdited      bool
}

type Category struct {
	Name string
}

// Given a title and text and a note type, creates a note object.
// Initialises dates and categories
func CreateNote(title string, text string, noteType string, categories []Category) *Note {
	note := Note{
		Title:         title,
		Text:          text,
		NoteType:      noteType,
		Categories:    categories,
		DatePublished: time.Now(),
		IsEdited:      false,
	}
	return &note
}

// Returns a formatted version of our note for printing out
func (note *Note) String() string {
	return fmt.Sprintf("Note titled %s from categories %s published on %d %s, %d",
		note.Title,
		note.Categories,
		note.DatePublished.Day(),
		note.DatePublished.Month(),
		note.DatePublished.Year())
}
