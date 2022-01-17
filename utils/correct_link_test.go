package utils

import (
	v1 "github.com/damiensedgwick/op/api/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnCorrectLink(t *testing.T) {
	assert.Equal(t, "https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html", ReturnCorrectLink("typescript"))
	assert.Equal(t, "https://reactjs.org/docs/getting-started.html", ReturnCorrectLink("react"))
}

func TestCheckLinkExists(t *testing.T) {
	incorrectLink, l1 := CheckLinkExists("billy-bob-bobo", v1.Links())
	assert.Equal(t, false, incorrectLink)
	assert.Equal(t, "", l1)

	correctLink, l2 := CheckLinkExists("react", v1.Links())
	assert.Equal(t, true, correctLink)
	assert.Equal(t, "https://reactjs.org/docs/getting-started.html", l2)
}
