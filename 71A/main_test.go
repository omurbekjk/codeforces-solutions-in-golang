package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, Solution("word"), "word", "it should return word")
	assert.Equal(t, Solution("localization"), "l10n", "it should return l10n")
	assert.Equal(t, Solution("internationalization"), "i18n", "it should return i18n")
	assert.Equal(t, Solution("pneumonoultramicroscopicsilicovolcanoconiosis"), "p43s", "it should return p43s")
	assert.Equal(t, Solution("omurbekkadyrbekov"), "o15v", "it should return o15v")
}
