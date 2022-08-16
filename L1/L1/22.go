package main

import (
	"fmt"
	"math/big"
)

func main22() {
	a := big.NewInt(5)
	b := big.NewInt(7)
	fmt.Println(subtraction(a, b))

}

func sum(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func division(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func multiplication(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func subtraction(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}
