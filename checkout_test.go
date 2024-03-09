package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanFunc(t *testing.T) {
	assert := assert.New(t)

	c := Checkout{Basket: make(map[string]int)}

	t.Run("Should return a map with correct counts and handle lower case entries when called", func(t *testing.T) {

		expected := map[string]int{"A": 1, "B":1}
		c.Scan("A")
		c.Scan("b")

		assert.Equal(c.Basket, expected)
		
	})

}

func TestGetTotalPriceFunc(t *testing.T) {
	assert := assert.New(t)

	c := Checkout {Basket: map[string]int { "A" : 6, "B": 3, "C": 2}}

	

	t.Run("Should return a total item count", func(t *testing.T) {

		itemCount :=c.GetTotalPrice()

		assert.Equal(itemCount, 11)
	})

}


