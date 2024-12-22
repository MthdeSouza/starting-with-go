package main

import (
	"fmt"
	"time"
)

func worker(workedId int, data chan int, delay int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workedId, x)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func main() {
	var workersQtd, timeDelay, processQtd int

	fmt.Print("Enter quantity of process necessaries: ")
	fmt.Scan(&processQtd)
	fmt.Print("Enter workers quantity: ")
	fmt.Scan(&workersQtd)
	fmt.Print("Enter delay to each process in seconds: ")
	fmt.Scan(&timeDelay)

	data := make(chan int)

	for i := 0; i < workersQtd; i++ {
		go worker(i, data, timeDelay)
	}

	for i := 0; i < processQtd; i++ {
		data <- i
	}

}
