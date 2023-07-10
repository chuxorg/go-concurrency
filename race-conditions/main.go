package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	raceCondition()
}

func raceCondition() {
	// when we run this program, we will get different results order.
	// You may see "Hello" or "World" printed first.
	// this is because the goroutines are racing to update the msg variable.
	///
	// Run go run --race . from the terminal in this directory (make sure raceCondition() is called in main)
	/*
		Here is the output on my machine:

		WARNING: DATA RACE
			Write at 0x0000011e2c50 by goroutine 7:
  			main.updateMsg()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:13 +0x6f
  			main.raceCondition.func1()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:25 +0x37

			Previous write at 0x0000011e2c50 by goroutine 8:
  			main.updateMsg()
      	    /Users/csailer/projects/go-concurrency/race-conditions/main.go:13 +0x6f
  			main.raceCondition.func2()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:26 +0x37

			Goroutine 7 (running) created at:
  			main.raceCondition()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:25 +0x44
  			main.main()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:17 +0x24

			Goroutine 8 (finished) created at:
  			main.raceCondition()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:26 +0x50
  			main.main()
      		/Users/csailer/projects/go-concurrency/race-conditions/main.go:17 +0x24
		==================
		The --race flag is great. It provides away to do a quick test to see if you are
		introducing a race condition which is a common bug in concurrent programs.	
	*/
	wg.Add(2)
	go updateMsg("Hello")
	go updateMsg("World")
	wg.Wait()
	fmt.Println(msg)
}