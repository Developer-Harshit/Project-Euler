/*
Number Letter Counts
Problem 17
*/
package main

import (
	"fmt"
)

var (
	onesMap = map[byte]string {
		'0' : "",
		'1' : "one",
		'2' : "two",
		'3' : "three",
		'4' : "four",
		'5' : "five",
		'6' : "six",
		'7' : "seven",
		'8' : "eight",
		'9' : "nine",
	}
	tensMap = map[byte]string {
		'0' : "",
		'1' : "ten",
		'2' : "twenty",
		'3' : "thirty",
		'4' : "forty",
		'5' : "fifty",
		'6' : "sixty",
		'7' : "seventy",
		'8' : "eighty",
		'9' : "ninety",

	}
	secondPlaceMap = map[byte]string {
		'0' : "ten",
		'1' : "eleven",
		'2' : "twelve",
		'3' : "thirteen",
		'4' : "fourteen",
		'5' : "fifteen",
		'6' : "sixteen",
		'7' : "seventeen",
		'8' : "eighteen",
		'9' : "nineteen",
	}
)
/* Converts number into its word form */
func countCharInNum(num int) int {
	
	if num < 1 {panic("WTF BRO")}
	counter := 0
	n := fmt.Sprintf("%d",num)
	if len(n) > 4 {panic("WTFF BROOO")}
	word := ""
	for i := len(n) - 1; i > -1; i-- {

		c := n[len(n) - i - 1]
		if c == '0' {continue}						// ONES PLACE
		if i == 0 {
			counter += len(onesMap[c])
			word += onesMap[c] + " "
		} else if i == 1 {								// TENS PLACE
			if c == '1' {
				counter += len(secondPlaceMap[n[len(n) - 1]])
				word += secondPlaceMap[n[len(n) - 1]] + " "
				break
			} else {
				counter += len(tensMap[c])
				word += tensMap[c] + " "
			}
		} else if i == 2 {								// HUNDREDTH PLACE
			counter += len(onesMap[c]) + len("hundred")
			word += onesMap[c] + " hundred "
			if n[len(n) - 1] != '0'|| n[len(n) - 2] != '0' {
				counter += len("and")
				word += "and "
			}
		} else if i == 3 {								// THOUSANDTH PLACE
			counter += len(onesMap[c]) + len("thousand")
			word += onesMap[c] + " thousand "
		}
	}
	return counter
}
func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += countCharInNum(i)
	}
	fmt.Println("Answer",sum)
}
