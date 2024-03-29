package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	number = (number * 2) + 100
	fmt.Println(number)
}
