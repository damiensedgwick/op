package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDNSearch(t *testing.T) {
	t1 := []string{
		"one",
		"two",
		"three",
	}
	t2 := []string{
		"cat",
		"dog",
		"fox",
	}
	assert.Equal(t, "https://developer.mozilla.org/en-US/search?q=one+two+three", MDNSearch(t1))
	assert.Equal(t, "https://developer.mozilla.org/en-US/search?q=cat+dog+fox", MDNSearch(t2))
}
