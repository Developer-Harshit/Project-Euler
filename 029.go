/*
Distinct Power
Problem 29
*/
package main

import (
	"fmt"
	"math/big"
)
func pow(base, exp int) string {
	bigNum := big.NewInt(1)
	for i := 0; i < exp; i++ {
		bigNum = bigNum.Mul(bigNum,big.NewInt(int64(base)))
	}
	return bigNum.String()
}
func main() {
	numMap := map[string]bool{}
	for a := 2; a < 101; a++ {
		for b := 2; b < 101; b++ {
			numMap[pow(a,b)] = true
		}
	}
	fmt.Println("Answer", len(numMap))
}
