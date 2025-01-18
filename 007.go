package main

import "fmt"

/*
<p>By listing the first six prime numbers: $2, 3, 5, 7, 11$, and $13$, we can see that the $6$th prime is $13$.</p>
<p>What is the $10\,001$st prime number?</p>
*/

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false 
		}
	}
	return true
}
func main() {
	
	var num,i int
	i = 10001 
//	i = 6
	for num = 2;i > 0; num++ {
		if isPrime(num) {
			i--
		}	
	}
	fmt.Println("Answer ",num-1)
}
