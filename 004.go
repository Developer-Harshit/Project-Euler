package main

import "fmt"

/*
<p>A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is $9009 = 91 \times 99$.</p>
<p>Find the largest palindrome made from the product of two $3$-digit numbers.</p>
*/

func isPalindrome(num int) bool {
	n := fmt.Sprint(num)
	a := 0
	b := len(n) - 1
	for a < b {
		if n[a] != n[b] {return false}
		a++
		b--
	}
	return true
}
func main() {
	largest := -1
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if isPalindrome(i*j){
				if largest < i*j {largest = i*j}		
			}
		}
	}
	fmt.Println("Largest",largest)
}
