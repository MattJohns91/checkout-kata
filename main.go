package main

import "strings"

type Checkout struct {
	Basket []string
}

type Prices map[string]int

type OfferPrice struct {
	NumberNeeded int
	DealPrice    int
}

//scan function to add item to basket
func (c *Checkout) Scan(item string) {

	prod := strings.ToUpper(item)

	c.Basket = append(c.Basket, prod)

}

func main() {

}
