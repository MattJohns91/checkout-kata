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

// func to return total price
func (c *Checkout) GetTotalPrice(p Prices, sPrice SpecialPrices) int {
	subTotal := 0

	for item, count := range c.Basket {
		// fmt.Printf(" \nItem is %v, count is %v \n", item, count)

		//check for special price available for item, then check special pricing threshold reached
		if specialPrice, ok := sPrice[item]; ok && count >= specialPrice.NumberNeeded {
			//add items at special price to subtotal
			subTotal += (count/ specialPrice.NumberNeeded) * specialPrice.DealPrice
			//add remaining items at standard price to subtotal
			subTotal += (count% specialPrice.NumberNeeded) * p[item]
		} else {
			subTotal += (p[item] * count)
		}
		
	}

	return subTotal
}

func main() {

	check := Checkout{Basket: make(map[string]int)}

	prices := Prices {
				"A": 50,
				"B": 30,
				"C": 20,
				"D": 15,
			}
		
	specialPrices := SpecialPrices{
				"A": {3, 130},
				"B": {2, 45},
			}

	check.Scan("A")
	check.Scan("A")
	check.Scan("B")
	check.Scan("C")
	check.Scan("D")
	check.Scan("B")
	check.Scan("B")

	total:= check.GetTotalPrice(prices, specialPrices)

	fmt.Printf("Total is %v \n", total)



}
