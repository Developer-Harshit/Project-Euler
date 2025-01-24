/*
Quadratic Primes
Problem 27
*/
package main

import "fmt"

var (
	cache = map[int]bool{}
)
func isPrime(num int) bool {
	if num < 0 {num*=-1}
	result,ok := cache[num]
	if ok {return result}
	for i := 2; i < num; i++ {
		if num % i == 0 {
			cache[num] = false
			return false 
		}
	}
	cache[num] = true
	return true
}
func CountPrimes(a, b int) int {
	n := 0
	for isPrime( n*n + n*a + b ) {
		n++
	}
	return n
}
func main() {
	largest := -1
	prod := -1
	for a := -999; a < 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			k := CountPrimes(a, b)
			if k >= largest {
				largest = k
				prod = a * b
			}
		}
	}
	fmt.Println("Answer",prod)
}
