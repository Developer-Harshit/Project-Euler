/*
Digit Power
Problem 30
*/
package main

import "fmt"
func isSumOfFifth(num int) bool {
	n := fmt.Sprintf("%d",num)
	sum := 0
	for _,c := range n {
		d := int(c - '0')
		sum += d * d * d * d * d
	}
	return sum == num
}
func main() {

	sum := 0
	for i := 2; i < 1000000; i++ {
		if isSumOfFifth(i) {sum += i}
	}
	fmt.Println("Answer",sum)
}

