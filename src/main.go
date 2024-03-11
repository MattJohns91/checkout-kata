package main

import (
	"fmt"
	"github.com/MattJohns91/checkout-kata/src/checkout"
)


func main() {

	check := checkout.Checkout{Basket: make(map[string]int)}

	prices := checkout.Prices {
				"A": 50,
				"B": 30,
				"C": 20,
				"D": 15,
			}
		
	specialPrices := checkout.SpecialPrices{
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
