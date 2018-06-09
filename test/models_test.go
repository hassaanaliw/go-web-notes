package test

import (
	"testing"
	"webnotes/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	category := model.Category{Name: "Hello",}
	assert.Equal(t, category.Name, "Hello", "Category name is not set correctly")
}

func TestCreateNote(t *testing.T) {
	note := model.Note{
		Title: "Test Note",
	}
	assert.Equal(t, note.Title, "Test Note", "Test Note not set correctly")
}
