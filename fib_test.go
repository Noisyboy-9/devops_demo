package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib_seedValues(t *testing.T) {
	assert.Equal(t, fib(0), 0)
	assert.Equal(t, fib(1), 1)
}

func TestFib_midValue(t *testing.T) {
	assert.Equal(t, fib(5), 5)
	assert.Equal(t, fib(6), 8)
}

func TestFib_bigValue(t *testing.T) {
	assert.Equal(t, fib(10), 55)
	assert.Equal(t, fib(11), 89)
}
