/*
Coin Sums
Problem 31
*/
package main

import "fmt"

/*
function dfs(amount, last) {
    if (amount === 0) return 1;
    const coins = [1, 2, 5, 10, 20, 50, 100, 200];
    let ans = 0;
    for (let x of coins) { // iterate over the coins
        if (amount- x >= 0 && x <= last) { // non-bigger
            ans += dfs(amount - x, x);  // recursive dfs
        }
    }
    return ans;
}

console.log(dfs(200, 200));
*/
func DFS(target, last int) int {
	if target == 0 {
		return 1
	}
	coins := []int {1,2,5,10,20,50,100,200}
	ans := 0

	for _,x := range coins {
		if target - x >= 0 && x <= last {
			ans += DFS(target - x, x)
		}
	}
	return ans
}
func main() {
	
	fmt.Println("Answer",DFS(200,200))
}
