package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	var index1, index2, index3, index4, index5, index6 int
	fmt.Scan(&index1, &index2, &index3, &index4, &index5, &index6)

	workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	workArray[index3], workArray[index4] = workArray[index4], workArray[index3]
	workArray[index5], workArray[index6] = workArray[index6], workArray[index5]

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}
