package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleSearch(t *testing.T) {
	t1 := []string{
		"golang",
	}
	t2 := []string{
		"react",
		"hooks",
	}
	assert.Equal(t, "https://www.google.com/search?q=golang", GoogleSearch(t1))
	assert.Equal(t, "https://www.google.com/search?q=react+hooks", GoogleSearch(t2))
}
