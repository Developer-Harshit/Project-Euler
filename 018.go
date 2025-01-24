/*
Maximum Path Sum I
Problem 18
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ID(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

var input string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

var (
	grid = map[string]int{}
	sx   = 0
	sy   = 0
)

func getNeighbour(x, y int) [][2]int {
	neighbour := [][2]int{}
	_, left := grid[ID(x-1, y+1)]
	if left {
		neighbour = append(neighbour, [2]int{x - 1, y + 1})
	}
	_, right := grid[ID(x+1, y+1)]
	if right {
		neighbour = append(neighbour, [2]int{x + 1, y + 1})
	}
	return neighbour
}

var routes int = 0

func DFS(cx, cy, s int) int {
	sum := grid[ID(cx, cy)] + s
	if cy == sy-1 {
		routes++
		return sum
	}
	maxSum := -1
	for _, u := range getNeighbour(cx, cy) {
		k := DFS(u[0], u[1], sum)
		if k > maxSum {
			maxSum = k
		}
	}
	return maxSum
}
func main() {

	for _, line := range strings.Split(input, "\n") {
		for j, n := range strings.Split(line, " ") {
			x := sx + j*2
			y := sy
			num, _ := strconv.Atoi(n)
			grid[ID(x, y)] = num
		}
		sx--
		sy++
	}
	// PRINTING GRID
	for j := 0; j < sy; j++ {
		for i := sx + 1; i < sx*-1; i++ {
			num, ok := grid[ID(i, j)]
			if !ok {
				fmt.Printf("   ")
			} else {
				fmt.Printf("%2d ", num)
			}
		}
		fmt.Println()
	}
	fmt.Printf("\nAnswer_> %d\tRoutes_> %d\n", DFS(0, 0, 0), routes)

}
