package main

import "fmt"

/*
<p>The sum of the squares of the first ten natural numbers is,</p>
$$1^2 + 2^2 + ... + 10^2 = 385.$$
<p>The square of the sum of the first ten natural numbers is,</p>
$$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
<p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
<p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
*/ 

func main() {

	SumOfSq := 0
	SqOfSum := 0
	for i := 100; i > 0; i-- {
		SumOfSq += i*i
		SqOfSum += i
	}
	SqOfSum = SqOfSum * SqOfSum
	fmt.Println("Answer ", SqOfSum - SumOfSq)
}
