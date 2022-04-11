package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceSpaces(t *testing.T) {
	t1 := []string{
		"one",
		"two",
		"three",
	}
	assert.Equal(t, "one-two-three", ReplaceSpaces(t1))

	t2 := []string{
		"oh",
		"my",
		"zsh",
	}
	assert.Equal(t, "oh-my-zsh", ReplaceSpaces(t2))
}
