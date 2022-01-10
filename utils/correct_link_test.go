package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnCorrectLink(t *testing.T) {
	assert.Equal(t, "https://developer.mozilla.org/en-US/docs/Web/JavaScript", ReturnCorrectLink("javascript"))
	assert.Equal(t, "https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html", ReturnCorrectLink("typescript"))
	assert.Equal(t, "https://reactjs.org/docs/getting-started.html", ReturnCorrectLink("react"))
}
