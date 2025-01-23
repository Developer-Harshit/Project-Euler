/*
Factorial Digit Sum
Problem 20
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {

	f := big.NewInt(1)
	for i := 1; i <= 100; i++ {	
	
		f.Mul(f, big.NewInt(int64(i)))
	}
	sum := 0
	for _,c := range f.String() {
		sum += int(c - '0')
	}
	fmt.Println("Answer",sum)
}
