package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	hour := d / 30
	minute := d % 30

	println("It is", hour, "hours", minute, "minutes")
}
