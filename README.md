# Checkout 
Checkout is a service emulating a basic checkout system written in Golang.

### Usage
`git clone git@github.com:MattJohns91/checkout-kata.git`
`cd src/`  
`go run main.go` for example use case.
`cd src/checkout`
`go test` to run unit tests.

1. Import the checkout package.
2. Create a new checkout instance (with empty basket, as seen in `main.go`).
3. Scan items using the Scan method.
4. Calculate total price using GetTotalPrice method, utilising desired pricing structure.

### Issues / Future Work
1. Currently items outside of the pricing structure can be added to the basket, potential for checks to be added to Scan method to verify addition against a list of known SKUs.
2. Potential for method to instantiate a new checkout with empty basket as a quality of life improvement.