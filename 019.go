/*
Counting Sundays
Problem 19 
*/
package main

import "fmt"

type Date struct{ 
	year int
	month int // 0-11
	day int   // 1-(27,28,30,31) max
}
func (d Date) MonthDays() int {
	switch d.month {
		case 1,3,5,7,8,10,12:
			return 31
		case 2:
			if d.year & 4 == 0 {return 29}
			return 28
		case 4,6,9,11:
			return 30
		default:
			return 0
	}
}
func (d *Date) AddMonth() {
	d.month++
	if d.month > 12 {
		d.month = 1
		d.year++
	}
}
func main() {
	weekday := 1 // Sunday = 0 to Saturday = 6
	curr := Date{year: 1901,day: 1,month: 1}
	counter := 0
	for curr.month != 12 || curr.year != 2000 {
		weekday += curr.MonthDays()
		weekday %= 7
		if weekday == 0 {counter++}
		curr.AddMonth()
	}
	fmt.Println("Answer",counter)
}
