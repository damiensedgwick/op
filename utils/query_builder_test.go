package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
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
	assert.Equal(t, "one+two+three", QueryBuilder(t1))
	assert.Equal(t, "cat+dog+fox", QueryBuilder(t2))
}
