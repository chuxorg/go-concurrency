package workgroups

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func workgroupExample() {
	var wg sync.WaitGroup

	words := []string{
		"Alpha", 
		"Beta", 
		"Gamma", 
		"Delta", 
		"Epsilon",
	}

	// Add the number of goroutines we are going to wait for (the length of the words slice)
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i + 1, x), &wg)
	} 
	wg.Wait()
	// if you comment out this line you will get a panic: sync: negative WaitGroup counter.
	// Why? Because the WaitGroup counter is 0 at this point and we are trying to subtract 1 from it.
    wg.Add(1)		
	printSomething("Something to print, again.", &wg)
}