package main

import (
	"errors"
	"fmt"
	"math"
)

func calcArea(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, errors.New("Radius must be a positive number")
	}

	pi := 3.14
	return pi * math.Pow(radius, 2), nil
}

func main() {
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	area, err := calcArea(radius)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The circle area is %g \n", area)
	}
}
