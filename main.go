package main

import "fmt"

func main() {
	var number float64
	fmt.Scan(&number)

	if number <= 0.0 {
		fmt.Printf("число %4.2f не подходит", number)
	} else if number > 10000 {
		fmt.Printf("%e", number)
	} else {
		number = number * number
		fmt.Printf("%.4f", number)
	}
}
