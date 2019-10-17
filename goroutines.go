package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	go sum(1, 2)
	time.Sleep(2 * time.Second)
	fmt.Println("Hello Main!")
}

func hello() {
	fmt.Println("Hello Goroutines!")
}

func sum(a int, b int) {
	fmt.Printf("Sum: %d\n", a+b)
}

/**
* * Channel
 */
func channel() {
	a := make(chan int)
	fmt.Printf("Type of a is %T", a)
}
