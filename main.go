package main

import "fmt"

func main() {
	var x, p, y int
	var year int = 0
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for x < y {
		x = x + (x*p)/100
		year++
	}
	fmt.Println(year)
}
