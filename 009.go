package main

import "fmt"

/*
<p>A Pythagorean triplet is a set of three natural numbers, $a \lt b \lt c$, for which,
$$a^2 + b^2 = c^2.$$</p>
<p>For example, $3^2 + 4^2 = 9 + 16 = 25 = 5^2$.</p>
<p>There exists exactly one Pythagorean triplet for which $a + b + c = 1000$.<br>Find the product $abc$.</p>
*/ 
func main() {
	
	for c := 1; c < 1000; c++ {
		for b := 1 ; b < c; b++ {
			for a := 1; a < b; a++ {
				
				if a+b+c == 1000 && a*a + b*b == c*c {
					fmt.Println(a*b*c)
					return
				}
			}	
		}
	}

}
