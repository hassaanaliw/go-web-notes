package test

import (
	"testing"
	"github.com/hassaanaliw/go-web-notes/model"
	"github.com/stretchr/testify/assert"
)

// Create a sample category
func TestCreateCategory(t *testing.T) {
	category := model.Category{Name: "Hello",}
	assert.Equal(t, category.Name, "Hello", "Category name is not set correctly")
}

// Create a Sample Note
func TestCreateNote(t *testing.T) {
	categories := []model.Category{}
	categories = append(categories, model.Category{
		Name: "Hi",
	})
	note := model.CreateNote("Test Note", "This is the test text", "Normal", categories)
	assert.Equal(t, note.Title, "Test Note", "Test Note not set correctly")
}
