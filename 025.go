/*
1000 Digit Fibonacci Number
Problem 25
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	
	index := 1
	a := big.NewInt(1)
	b := big.NewInt(1)
	for len(a.String()) < 1000 {
		
		c := new(big.Int).Set(b)
		b.Add(a, b)
		a.Set(c)
		index++
	}
	fmt.Println("Answer",index)
}
