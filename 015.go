/*
Lattice Path
Problem 15
*/

package main

import "fmt"

var (
	offsets = [][2]int{{1, 0}, {0, 1}}
	cache   = map[int]int{}
	n       = 20
)

func ID(x, y int) int {
	return x + y*(n+1)
}

func DFS(x, y int) int {
	if x < 0 || y < 0 || x > n || y > n {
		return 0
	}
	if x == n && y == n {
		return 1
	}
	c, ok := cache[ID(x, y)]
	if ok { return c }
	count := 0
	for _, o := range offsets {
		count += DFS(x + o[0], y + o[1])
	}
	cache[ID(x, y)] = count
	return count
}
func main() {
	fmt.Println("Answer", DFS(0, 0))
}
