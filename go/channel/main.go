package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int, i int) {
	value := rand.Intn(10)
	fmt.Printf("%d - Calculated Random Value: %d\n", i, value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Printf("%d - Only Executes after another goroutine performs a receive on the channel\n", i)
}

func main() {
	fmt.Println("Go Channel Tutorial")

	valueChannel := make(chan int, 10)
	defer close(valueChannel)

	for i := 1; i <= 10; i++ {
		go CalculateValue(valueChannel, i)
	}

	for value := range valueChannel {
		fmt.Println(value)
	}
}
