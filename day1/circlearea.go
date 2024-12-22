package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	pi := 3.14

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	area := pi * math.Pow(float64(radius), 2)

	fmt.Printf("The circle area is %g \n", area)
}
