package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string 
	Amount int 
}


func main(){
	//raceConditionBankBalance() // no mutex lock and unlock (produces a race condition)
	threadSafeBalance() // mutex lock and unlock (runs correctly)
}

func raceConditionBankBalance(){

	// var for Bank balance
	var bankBalance int
	// print out starting Bank Balance
	fmt.Printf("Initial Bank Balance: %d\n", bankBalance)
	// define weekly salary
	incomes := []Income{
		{Source: "Salary", Amount: 1000},
		{Source: "Freelance", Amount: 100},
		{Source: "Side Hustle", Amount: 50},
		{Source: "Investments", Amount: 10},
	}
	wg.Add(len(incomes))

	// loop through 52 weeks of the year
    for i, income := range incomes {
		
		go func(i int, income Income) {
			defer wg.Done()	
				for week := 1; week <= 52; week++ {
					temp:= bankBalance
					temp += income.Amount
					bankBalance = temp
					fmt.Printf("Week %d - %s: $%d\n", week, income.Source, bankBalance)
				}
		
			}(i, income)	
		
	}
	wg.Wait()
	// print out final bankBalance
	fmt.Printf("Final bankBalance: $%d\n", bankBalance)
}

func threadSafeBalance(){
	// var for Bank balance
	var bankBalance int
	// var for mutex
	var mutex sync.Mutex

	// print out starting Bank Balance
	fmt.Printf("Initial Bank Balance: %d\n", bankBalance)
	// define weekly salary
	incomes := []Income{
		{Source: "Salary", Amount: 1000},
		{Source: "Freelance", Amount: 100},
		{Source: "Side Hustle", Amount: 50},
		{Source: "Investments", Amount: 10},
	}
	wg.Add(len(incomes))

	// loop through 52 weeks of the year
    for i, income := range incomes {
		
		go func(i int, income Income) {
			defer wg.Done()	
				for week := 1; week <= 52; week++ {
					mutex.Lock()
					temp := bankBalance
					temp += income.Amount
					bankBalance = temp
					fmt.Printf("Week %d - %s: $%d\n", week, income.Source, bankBalance)
					mutex.Unlock()
				}
		
			}(i, income)	
		
	}
	wg.Wait()
	// print out final bankBalance
	fmt.Printf("Final bankBalance: $%d\n", bankBalance)
}