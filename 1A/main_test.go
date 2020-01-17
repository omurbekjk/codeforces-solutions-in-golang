package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, Solution(6, 6, 4), int64(4), "should return 4")
	assert.Equal(t, Solution(5, 5, 5), int64(1), "should return 1")
	assert.Equal(t, Solution(5, 5, 25), int64(1), "should return 1")
	assert.Equal(t, Solution(12, 12, 1), int64(144), "should return 144")
	assert.Equal(t, Solution(3, 3, 2), int64(4), "should return 4")
	assert.Equal(t, Solution(1000000000, 1000000000, 1), int64(1000000000000000000), "should return 1000000000000000000")
}
