/*
Longest Collatz Sequence
Problem 14
*/
package main

import "fmt"


func nextSequence(n int) int {
	if n % 2 == 0 {return n/2}
	return 3*n + 1
}
func main() {
	k := 1000000
	largest := -1
	maxCount := 0
	for i := 1; i < k; i++ {
		n := i
		count := 1
		for n > 1 {
			n = nextSequence(n)
			count++
		}
		if count > maxCount {maxCount = count;largest = i}
	}
	fmt.Println("Answer",largest)
}
