package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, Solution(8), "YES", "answer should be YES")
	assert.Equal(t, Solution(1), "NO", "answer should be NO")
	assert.Equal(t, Solution(9), "NO", "answer should be NO")
	assert.Equal(t, Solution(100), "YES", "answer should be YES")
	assert.Equal(t, Solution(17), "NO", "answer should be NO")
}
