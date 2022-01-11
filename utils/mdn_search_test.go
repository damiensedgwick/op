package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMDNSearch(t *testing.T) {
	assert.Equal(t, "https://developer.mozilla.org/en-US/search?q=div", MDNSearch("div"))
	assert.Equal(t, "https://developer.mozilla.org/en-US/search?q=p", MDNSearch("p"))
	assert.Equal(t, "https://developer.mozilla.org/en-US/search?q=table", MDNSearch("table"))
}
