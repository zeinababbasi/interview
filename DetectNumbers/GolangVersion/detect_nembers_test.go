package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectNumber(t *testing.T) {
	a := assert.New(t)

	t.Run("numbers", func(t *testing.T) {
		a.True(DetectNumber("10"))
		a.True(DetectNumber("-10"))
		a.True(DetectNumber("10.1"))
		a.True(DetectNumber("-10.1"))
		a.True(DetectNumber("1e5"))
	})

	t.Run("non numbers", func(t *testing.T) {
		a.False(DetectNumber("a"))
		a.False(DetectNumber("x 1"))
		a.False(DetectNumber("a -2"))
		a.False(DetectNumber("-"))
	})
}
