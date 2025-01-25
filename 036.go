package main

import (
	"fmt"
	"math/big"
)

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {return false}
		left++
		right--
	}
	return true
}
func main() {
	counter := int64(0)
	for i := int64(1); i < 1000000; i++ {
		b := big.NewInt(i)
		if isPalindrome(b.Text(10)) && isPalindrome(b.Text(2)) {
			counter+=i;
		}
	}
	fmt.Println("Answer",counter)
}
