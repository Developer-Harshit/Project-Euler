package main

import "fmt"

/*
<p>The prime factors of $13195$ are $5, 7, 13$ and $29$.</p>
<p>What is the largest prime factor of the number $600851475143$?</p>
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
	
	largest := 1
	num := 600851475143;
//	num = 13195 
	for i := 2; num != 1; i++ {
		
		if !isPrime(i) {continue}
		largest = i
		// divide until num is not divisible anymore
		for num % i == 0 {
			num /= i
//			fmt.Printf("Dividing by %d \t Num Becomes %d\n",i,num)		
		}

	}
	fmt.Println(largest)

}
