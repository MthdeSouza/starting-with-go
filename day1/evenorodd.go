package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num1)

	remainder := num1 % 2
	result := ""

	if remainder == 0 {
		result = "even"
	} else {
		result = "odd"
	}

	fmt.Printf("The number %d is %s\n", num1, result)
}
