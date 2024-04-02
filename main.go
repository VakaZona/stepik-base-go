package main

import "fmt"

func main() {
	var n int
	var number int
	var sum int = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if (number%8 == 0) && (number >= 10) && (number <= 99) {
			sum += number
		}
	}
	fmt.Println(sum)
}
