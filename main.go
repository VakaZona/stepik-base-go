package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	hour := d / 30
	minute := (d % 30) * 2

	fmt.Println("It is", hour, "hours", minute, "minutes.")
}
