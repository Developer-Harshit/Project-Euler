/*
Power Digit Sum
Problem 16
*/
package main

import (
	"fmt"
	"math/big"
)
func main() {

	str := "1"
	p := 1000
	for i := 0; i < p; i++ {
		str += "0"
	}
	bigNum,_ := new(big.Int).SetString(str, 2)
	sum := 0
	for _,d := range bigNum.String() {
		sum += int(d - '0')
	}

	fmt.Println("Answer",sum)
}
