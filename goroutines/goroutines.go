package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A goroutine is a function that is capable of running concurrently with other functions.
// To create a goroutine, use the keyword "go" followed by the name of the func being called
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
// The first goroutine is implicit and is the main function 
func main() {
	for i := 0; i < 10; i++ {
		// second goroutine
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}