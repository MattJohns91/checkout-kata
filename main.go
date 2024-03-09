package main

import (
	"fmt"
	"strings"
)

type Checkout struct {
	Basket map[string]int
}

// standard pricing
type Prices map[string]int

// offer pricing map
type SpecialPrices map[string]Offer

type Offer struct {
	NumberNeeded int
	DealPrice    int
}

// scan function to add item to basket
func (c *Checkout) Scan(item string) {

	//handle lowercase
	prod := strings.ToUpper(item)

	c.Basket[prod]++

}

//func to return total price
func (c *Checkout) GetTotalPrice(p Prices) int {
	subTotal := 0

	for item, count := range c.Basket {
		fmt.Printf(" \nItem is %v, count is %v \n", item, count)
		subTotal += (p[item] * count)
	}

	return subTotal
}



func main() {

	check := Checkout {Basket: make (map[string]int) }


	check.Scan("A")
	check.Scan("A")
	check.Scan("B")
	check.Scan("C")
	check.Scan("D")
	check.Scan("B")
	check.Scan("B")



	




}
