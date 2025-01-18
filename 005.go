package main

import "fmt"

/*
<p>$2520$ is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.</p>
<p>What is the smallest positive number that is <strong class="tooltip">evenly divisible<span class="tooltiptext">divisible with no remainder</span></strong> by all of the numbers from $1$ to $20$?</p>
*/

func isDivisibleByAll(num, n int) bool {
	for i := 2; i < n; i++ {
		if num % i != 0 {return false}
	}
	return true
}
func main() {
	
	start 	:= 1*2*3*5*7*11*13*17*19 // ofcourse the number must be divisible by these primes
	inc		:= 19 // incrementing by largest prime
	for num := start;;num+= inc {
		if isDivisibleByAll(num, 20) {
			fmt.Print("Answer ",num)
			break
		}
	}

}
