package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanFunc(t *testing.T) {
	assert := assert.New(t)

	c := Checkout{}

	t.Run("Should return a slice with length correct length and handle lower case entries when called", func(t *testing.T) {

		expected := []string{"A","B"}
		c.Scan("A")
		c.Scan("b")

		assert.Equal(c.Basket, expected)
		assert.Len(c.Basket, 2)
	})

}
