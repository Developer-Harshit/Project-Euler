/*
<p>The sum of the primes below $10$ is $2 + 3 + 5 + 7 = 17$.</p>
<p>Find the sum of all the primes below two million.</p>
*/
package main

import "fmt"

var (
	primes = []int{}
	largest = 2
)
func isPrime(num int) bool {
	for _,i := range primes {
		if num % i == 0 {return false}
	}
	for i := largest; i < num; i++ {
		if num % i == 0 {return false}
	}
	return true
}
func main() {
	
	sum := 0
	n	:= 2000000 // two million
	for i := 2; i < n; i++ {
		if isPrime(i) {
			sum += i
			largest = i
			{primes = append(primes, i)}
		}
	}
	fmt.Println("answer",sum)
}

