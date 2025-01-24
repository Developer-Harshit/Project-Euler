/*
Number Spiral Diagonals
Problem 28
*/

package main

import "fmt"
func main() {
	i := 1
	n := 1001
	sum := 1
	for r := 1; r <= int(n/2); r++ {
		for a:=0;a<4;a++ {
			i += r * 2
			sum += i
		}
	}
	fmt.Println("Answer _>",sum)
}
