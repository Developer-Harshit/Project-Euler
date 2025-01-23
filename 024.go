/*
Lexicographic Permutations
Problem 24
*/
package main

import (
	"fmt"
	"slices"
)


func nextPermutation(str string) (string,bool) {
	a := []byte(str[:])
	size := len(a)
	var i, j int
	for i = size - 2; i > -1 && !(a[i] < a[i+1]); i-- {}
	if i == -1 {
		return str,false
	}

	for j = size - 1; j > i && !(a[i] < a[j]); j-- {}
//	if j == i {return str,false}
	
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp


	subSlice := a[i+1:]
	slices.Reverse(subSlice)
	a = append(a[:i+1], subSlice...) 

	return string(a[:]),true
}
func main() {
	
	shouldRun := true
	order := "0123456789"
	n := 1000000 // ONE MILLION
	for i:=1;shouldRun && i != n;i++ {
		order,shouldRun = nextPermutation(order)
	}
		fmt.Printf("Index: %4d  Order: %s  Bool: %v\n",n,order,shouldRun)
}
