package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if 0 != number/10000 {
		fmt.Print(number / 10000)
	} else if 0 != number/1000 {
		fmt.Println(number / 1000)
	} else if 0 != number/100 {
		fmt.Println(number / 100)
	} else if 0 != number/10 {
		fmt.Println(number / 10)
	} else {
		fmt.Println(number)
	}
}
