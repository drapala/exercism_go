package main

import (
	"fmt"
	"math/big"
)

// Calculates g^a
func PowerBig(g, a *big.Int) *big.Int {
	return new(big.Int).Exp(g, a, nil)
}

func main() {
	fmt.Println(PowerBig(big.NewInt(2), big.NewInt(3)))
}