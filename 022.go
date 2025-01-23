/*
Names Scores
Problem 22
*/
package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
)


func main() {
	
	inputPath := "data/names.txt"
	data, err := os.ReadFile(inputPath)
	if err != nil {panic(err)}
	str := string(data)
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\"", "")
	s := strings.Split(str, ",")

	sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})
	scoreSum := 0
	for i,val := range s {
		val = strings.Trim(val, "\n")
		score := 0
		for _,c := range val {
			score += int(c - 'A') + 1
		}

		fmt.Println(score,val,big.NewInt(int64('A'-'A')+1))
		score *= i + 1
		scoreSum += score 
	}
	fmt.Println("Answer",scoreSum)
}
