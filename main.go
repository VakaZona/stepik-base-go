package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	number1 := number / 100
	number2 := (number % 100) / 10
	number3 := number % 10

	if (number1 == number2) || (number1 == number3) || (number2 == number3) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
