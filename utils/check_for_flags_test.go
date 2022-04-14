package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidFlag(t *testing.T) {
	t1 := "-g"
	t2 := "banana"
	assert.Equal(t, true, IsValidFlag(t1))
	assert.Equal(t, false, IsValidFlag(t2))
}
