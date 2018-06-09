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
