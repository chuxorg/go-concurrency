package raceconditions

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func raceConditionExample() {
	// run: go run --race . from the terminal in this directory
	raceCondition()
}


func updateMsg(s string) {
	defer wg.Done()
	msg = s
}

func raceCondition() {
	// when we run this program, we will get different results order.
	// You may see "Hello" or "World" printed first.
	// this is because the goroutines are racing to update the msg variable.

	wg.Add(2)
	go updateMsg("Hello")
	
	go updateMsg("World")
	wg.Wait()
	fmt.Println(msg)
}

