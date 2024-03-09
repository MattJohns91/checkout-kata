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

func TestGetTotalCountFunc(t *testing.T) {
	assert := assert.New(t)

	c := Checkout {Basket: map[string]int { "A" : 6, "B": 3, "C": 0}}

// 	testP := Prices{
// 	"A" : 50,
// 	"B" : 120,
// 	"C" : 3,
// 	"D" : 12,
// }

	

	t.Run("Should return a total price using standard pricing", func(t *testing.T) {

		totalPrice :=c.GetTotalPrice()

		assert.Equal(totalPrice,660)
	})

}


