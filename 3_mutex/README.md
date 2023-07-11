# Mutex

This first function is :```raceConditionBankBalance```. Calculates the the amount from all income sources and gives a total income for each week of the year and then sums all weeks
for an anual income value. 

We're setting up Workgroups for each source then running a ```goroutine``` for each Income Source so we can calculate these totals. There is a bug. The bug is a race condition that occurs when setting the value of ```bankBalance```. The bug happens because bankBalance is a *Shared Resource* among the ```goroutines```. Each
routine will "race" to calculate the balances and there will only be one winner. We will only see the value of
```bankBalance``` set by the winning ```goroutine```



```go
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
```

Make sure you are in the ```3_mutex``` dir and run 
```bash
 $ go run --race .
```
You will see output similar to the following:


```go
base) Chucks-MacBook-Pro:3_mutex csailer$ go run --race .
Initial Bank Balance: 0
Week 1 - Freelance: $1100
Week 1 - Investments: $1110
Week 2 - Investments: $1120
Week 2 - Freelance: $1220
Week 3 - Freelance: $1320
Week 4 - Freelance: $1420
Week 5 - Freelance: $1530
Week 3 - Investments: $1430
Week 6 - Freelance: $1630
Week 7 - Freelance: $1740
Week 4 - Investments: $1640
Week 5 - Investments: $1850
Week 8 - Freelance: $1840
Week 6 - Investments: $1860
Week 7 - Investments: $1970
Week 9 - Freelance: $1960
Week 8 - Investments: $1980
Week 10 - Freelance: $2090
==================
Week 11 - Freelance: $2190
WARNING: DATA RACE    
Read at 0x00c00012c018 by goroutine 9:
  main.raceConditionBankBalance.func1()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:42 +0xd0
  main.raceConditionBankBalance.func2()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:48 +0x74

Previous write at 0x00c00012c018 by goroutine 7:
  main.raceConditionBankBalance.func1()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:44 +0xe5
  main.raceConditionBankBalance.func2()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:48 +0x74

Goroutine 9 (running) created at:
Week 12 - Freelance: $2290
  main.raceConditionBankBalance()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:39 +0x329
Week 9 - Investments: $1990
  main.main()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:17 +0x24

Goroutine 7 (running) created at:
  main.raceConditionBankBalance()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:39 +0x329
  main.main()
      /your/path/to/your/code/go-concurrency/3_mutex/main.go:17 +0x24
==================
Week 1 - Salary: $1000
Week 1 - Side Hustle: $2450
Week 2 - Salary: $3450
Week 3 - Salary: $4500
Week 2 - Side Hustle: $3500
Week 3 - Side Hustle: $5550
Week 4 - Side Hustle: $5600
Week 5 - Side Hustle: $5650
Week 10 - Investments: $2400
Week 11 - Investments: $5710
Week 12 - Investments: $5720
Week 6 - Side Hustle: $5700
Week 4 - Salary: $5500
Week 7 - Side Hustle: $5780
Week 5 - Salary: $6780
Week 13 - Investments: $5730
Week 14 - Investments: $7840
Week 15 - Investments: $7850
Week 16 - Investments: $7860
Week 8 - Side Hustle: $6830
Week 9 - Side Hustle: $7920
Week 10 - Side Hustle: $7970
Week 11 - Side Hustle: $8020
Week 12 - Side Hustle: $8070
Week 13 - Freelance: $2390
Week 14 - Freelance: $8220
Week 13 - Side Hustle: $8120
Week 17 - Investments: $7870
Week 14 - Side Hustle: $8370
Week 18 - Investments: $8380
Week 15 - Side Hustle: $8430
Week 19 - Investments: $8440
Week 20 - Investments: $8500
Week 21 - Investments: $8510
Week 22 - Investments: $8520
Week 23 - Investments: $8530
Week 24 - Investments: $8540
Week 25 - Investments: $8550
Week 26 - Investments: $8560
Week 16 - Side Hustle: $8490
Week 17 - Side Hustle: $8620
Week 18 - Side Hustle: $8670
Week 19 - Side Hustle: $8720
Week 15 - Freelance: $8320
Week 16 - Freelance: $8870
Week 17 - Freelance: $8970
Week 6 - Salary: $7830
Week 18 - Freelance: $9070
Week 7 - Salary: $10070
Week 20 - Side Hustle: $8770
Week 27 - Investments: $8570
Week 21 - Side Hustle: $11230
Week 28 - Investments: $11180
Week 29 - Investments: $11290
Week 8 - Salary: $11170
Week 30 - Investments: $11300
Week 9 - Salary: $12300
Week 31 - Investments: $12310
Week 10 - Salary: $13310
Week 32 - Investments: $13320
Week 33 - Investments: $14330
Week 34 - Investments: $14340
Week 35 - Investments: $14350
Week 36 - Investments: $14360
Week 11 - Salary: $14320
Week 37 - Investments: $14370
Week 19 - Freelance: $10170
Week 12 - Salary: $15380
Week 20 - Freelance: $15480
Week 38 - Investments: $14380
Week 39 - Investments: $16590
Week 40 - Investments: $16600
Week 41 - Investments: $16610
Week 42 - Investments: $16620
Week 43 - Investments: $16630
Week 21 - Freelance: $16580
Week 22 - Side Hustle: $11280
Week 23 - Side Hustle: $16790
Week 22 - Freelance: $16740
Week 44 - Investments: $16640
Week 45 - Investments: $16950
Week 23 - Freelance: $16940
Week 46 - Investments: $16960
Week 24 - Freelance: $17060
Week 24 - Side Hustle: $16840
Week 25 - Freelance: $17170
Week 25 - Side Hustle: $17220
Week 13 - Salary: $16480
Week 14 - Salary: $18370
Week 26 - Side Hustle: $17370
Week 27 - Side Hustle: $19420
Week 28 - Side Hustle: $19470
Week 29 - Side Hustle: $19520
Week 15 - Salary: $19370
Week 16 - Salary: $20570
Week 17 - Salary: $21570
Week 30 - Side Hustle: $19570
Week 31 - Side Hustle: $22620
Week 26 - Freelance: $17320
Week 32 - Side Hustle: $22670
Week 33 - Side Hustle: $22820
Week 34 - Side Hustle: $22870
Week 18 - Salary: $22570
Week 27 - Freelance: $22770
Week 28 - Freelance: $24020
Week 19 - Salary: $24020
Week 20 - Salary: $25120
Week 21 - Salary: $26120
Week 22 - Salary: $27120
Week 35 - Side Hustle: $22920
Week 36 - Side Hustle: $28170
Week 37 - Side Hustle: $28220
Week 29 - Freelance: $24120
Week 30 - Freelance: $28370
Week 38 - Side Hustle: $28270
Week 47 - Investments: $17070
Week 48 - Investments: $28530
Week 49 - Investments: $28540
Week 50 - Investments: $28550
Week 23 - Salary: $28120
Week 39 - Side Hustle: $28520
Week 24 - Salary: $29560
Week 40 - Side Hustle: $29610
Week 25 - Salary: $30610
Week 41 - Side Hustle: $30660
Week 26 - Salary: $31660
Week 42 - Side Hustle: $31710
Week 51 - Investments: $28560
Week 43 - Side Hustle: $31760
Week 52 - Investments: $31770
Week 44 - Side Hustle: $31820
Week 31 - Freelance: $28470
Week 32 - Freelance: $31970
Week 33 - Freelance: $32070
Week 34 - Freelance: $32170
Week 35 - Freelance: $32270
Week 36 - Freelance: $32370
Week 37 - Freelance: $32470
Week 38 - Freelance: $32570
Week 39 - Freelance: $32670
Week 40 - Freelance: $32770
Week 41 - Freelance: $32870
Week 42 - Freelance: $32970
Week 43 - Freelance: $33070
Week 27 - Salary: $32660
Week 28 - Salary: $34170
Week 45 - Side Hustle: $31870
Week 29 - Salary: $35170
Week 30 - Salary: $36220
Week 31 - Salary: $37220
Week 32 - Salary: $38220
Week 33 - Salary: $39220
Week 44 - Freelance: $33170
Week 46 - Side Hustle: $35220
Week 34 - Salary: $40220
Week 35 - Salary: $41370
Week 36 - Salary: $42370
Week 47 - Side Hustle: $40370
Week 48 - Side Hustle: $43420
Week 49 - Side Hustle: $43470
Week 45 - Freelance: $40320
Week 46 - Freelance: $43620
Week 47 - Freelance: $43720
Week 48 - Freelance: $43820
Week 49 - Freelance: $43920
Week 50 - Freelance: $44020
Week 51 - Freelance: $44120
Week 37 - Salary: $43370
Week 52 - Freelance: $44220
Week 50 - Side Hustle: $43520
Week 38 - Salary: $45220
Week 39 - Salary: $46270
Week 40 - Salary: $47270
Week 51 - Side Hustle: $45270
Week 52 - Side Hustle: $48320
Week 41 - Salary: $48270
Week 42 - Salary: $49320
Week 43 - Salary: $50320
Week 44 - Salary: $51320
Week 45 - Salary: $52320
Week 46 - Salary: $53320
Week 47 - Salary: $54320
Week 48 - Salary: $55320
Week 49 - Salary: $56320
Week 50 - Salary: $57320
Week 51 - Salary: $58320
Week 52 - Salary: $59320
Final bankBalance: $59320
Found 1 data race(s)
exit status 66
```
In the above output, notice what is going on. the ```--race``` flag is detecting race conditions when setting the value of the shared variable ```bankBalance```. These types
of errors are nasty. Its easy to see them with go's ```--race``` switch but without it, race conditions make finding bugs tricky, at best.

In the below example, the race condition is fixed by using a ```Mutex``` to acquire and release a *Lock* on the shared variable ```bankBalance```. A *lock* is a way to ensure only one thread at a time can access a shared resource such as a variable. When the lock is aquired, the lock ensures that all other threads are denied access to the shared resource until the lock is released by calling *Unlock*. 

```go
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
		defer wg.Done()	
		go func(i int, income Income) {
		        
				for week := 1; week <= 52; week++ {
					mutex.Lock()
					temp := bankBalance
					temp += income.Amount
					bankBalance = temp
					mutex.Unlock()
					fmt.Printf("Week $%d - %s: %d\n", week, income.Source, bankBalance)
				}
		
			}(i, income)	
		
	}
	wg.Wait()
	// print out final bankBalance
	fmt.Printf("Final bankBalance: $%d\n", bankBalance)
}
```
Return to the ```main.go``` file, uncomment ```threadSafeBalance()``` in the ```main()``` function and run:  

```bash
$ go run --race .
```

Now, the race condition is gone and the function works as expected all thanks to the lock and unlock methods of the ```Mutex```

```shell
$ go run --race . 

Week 1 - Salary: $1000
Week 2 - Salary: $2000
Week 3 - Salary: $3000
Week 4 - Salary: $4000
Week 5 - Salary: $5000
Week 6 - Salary: $6000
Week 7 - Salary: $7000
Week 8 - Salary: $8000
Week 9 - Salary: $9000
Week 10 - Salary: $10000
Week 11 - Salary: $11000
Week 12 - Salary: $12000
Week 13 - Salary: $13000
Week 14 - Salary: $14000
Week 15 - Salary: $15000
Week 16 - Salary: $16000
Week 17 - Salary: $17000
Week 18 - Salary: $18000
Week 19 - Salary: $19000
Week 20 - Salary: $20000
Week 21 - Salary: $21000
Week 22 - Salary: $22000
Week 23 - Salary: $23000
Week 24 - Salary: $24000
Week 25 - Salary: $25000
Week 26 - Salary: $26000
Week 27 - Salary: $27000
Week 28 - Salary: $28000
Week 29 - Salary: $29000
Week 30 - Salary: $30000
Week 31 - Salary: $31000
Week 32 - Salary: $32000
Week 33 - Salary: $33000
Week 34 - Salary: $34000
Week 35 - Salary: $35000
Week 36 - Salary: $36000
Week 37 - Salary: $37000
Week 38 - Salary: $38000
Week 39 - Salary: $39000
Week 40 - Salary: $40000
Week 41 - Salary: $41000
Week 42 - Salary: $42000
Week 43 - Salary: $43000
Week 44 - Salary: $44000
Week 1 - Side Hustle: $44050
Week 2 - Side Hustle: $44100
Week 3 - Side Hustle: $44150
Week 4 - Side Hustle: $44200
Week 5 - Side Hustle: $44250
Week 6 - Side Hustle: $44300
Week 7 - Side Hustle: $44350
Week 8 - Side Hustle: $44400
Week 9 - Side Hustle: $44450
Week 10 - Side Hustle: $44500
Week 11 - Side Hustle: $44550
Week 12 - Side Hustle: $44600
Week 13 - Side Hustle: $44650
Week 14 - Side Hustle: $44700
Week 15 - Side Hustle: $44750
Week 16 - Side Hustle: $44800
Week 17 - Side Hustle: $44850
Week 18 - Side Hustle: $44900
Week 19 - Side Hustle: $44950
Week 20 - Side Hustle: $45000
Week 21 - Side Hustle: $45050
Week 22 - Side Hustle: $45100
Week 23 - Side Hustle: $45150
Week 24 - Side Hustle: $45200
Week 25 - Side Hustle: $45250
Week 26 - Side Hustle: $45300
Week 27 - Side Hustle: $45350
Week 28 - Side Hustle: $45400
Week 29 - Side Hustle: $45450
Week 30 - Side Hustle: $45500
Week 31 - Side Hustle: $45550
Week 32 - Side Hustle: $45600
Week 33 - Side Hustle: $45650
Week 34 - Side Hustle: $45700
Week 1 - Investments: $45710
Week 1 - Freelance: $45810
Week 45 - Salary: $46810
Week 2 - Investments: $46820
Week 3 - Investments: $46830
Week 4 - Investments: $46840
Week 5 - Investments: $46850
Week 6 - Investments: $46860
Week 7 - Investments: $46870
Week 8 - Investments: $46880
Week 9 - Investments: $46890
Week 10 - Investments: $46900
Week 11 - Investments: $46910
Week 12 - Investments: $46920
Week 13 - Investments: $46930
Week 14 - Investments: $46940
Week 15 - Investments: $46950
Week 16 - Investments: $46960
Week 17 - Investments: $46970
Week 18 - Investments: $46980
Week 19 - Investments: $46990
Week 20 - Investments: $47000
Week 21 - Investments: $47010
Week 22 - Investments: $47020
Week 23 - Investments: $47030
Week 24 - Investments: $47040
Week 25 - Investments: $47050
Week 26 - Investments: $47060
Week 27 - Investments: $47070
Week 28 - Investments: $47080
Week 29 - Investments: $47090
Week 30 - Investments: $47100
Week 31 - Investments: $47110
Week 32 - Investments: $47120
Week 33 - Investments: $47130
Week 34 - Investments: $47140
Week 35 - Investments: $47150
Week 36 - Investments: $47160
Week 37 - Investments: $47170
Week 38 - Investments: $47180
Week 39 - Investments: $47190
Week 40 - Investments: $47200
Week 41 - Investments: $47210
Week 42 - Investments: $47220
Week 43 - Investments: $47230
Week 44 - Investments: $47240
Week 45 - Investments: $47250
Week 46 - Investments: $47260
Week 47 - Investments: $47270
Week 48 - Investments: $47280
Week 49 - Investments: $47290
Week 50 - Investments: $47300
Week 51 - Investments: $47310
Week 52 - Investments: $47320
Week 46 - Salary: $48320
Week 47 - Salary: $49320
Week 48 - Salary: $50320
Week 49 - Salary: $51320
Week 50 - Salary: $52320
Week 51 - Salary: $53320
Week 52 - Salary: $54320
Week 2 - Freelance: $54420
Week 3 - Freelance: $54520
Week 4 - Freelance: $54620
Week 5 - Freelance: $54720
Week 6 - Freelance: $54820
Week 7 - Freelance: $54920
Week 8 - Freelance: $55020
Week 35 - Side Hustle: $55070
Week 36 - Side Hustle: $55120
Week 37 - Side Hustle: $55170
Week 38 - Side Hustle: $55220
Week 39 - Side Hustle: $55270
Week 40 - Side Hustle: $55320
Week 41 - Side Hustle: $55370
Week 42 - Side Hustle: $55420
Week 43 - Side Hustle: $55470
Week 44 - Side Hustle: $55520
Week 45 - Side Hustle: $55570
Week 46 - Side Hustle: $55620
Week 47 - Side Hustle: $55670
Week 48 - Side Hustle: $55720
Week 49 - Side Hustle: $55770
Week 50 - Side Hustle: $55820
Week 51 - Side Hustle: $55870
Week 52 - Side Hustle: $55920
Week 9 - Freelance: $56020
Week 10 - Freelance: $56120
Week 11 - Freelance: $56220
Week 12 - Freelance: $56320
Week 13 - Freelance: $56420
Week 14 - Freelance: $56520
Week 15 - Freelance: $56620
Week 16 - Freelance: $56720
Week 17 - Freelance: $56820
Week 18 - Freelance: $56920
Week 19 - Freelance: $57020
Week 20 - Freelance: $57120
Week 21 - Freelance: $57220
Week 22 - Freelance: $57320
Week 23 - Freelance: $57420
Week 24 - Freelance: $57520
Week 25 - Freelance: $57620
Week 26 - Freelance: $57720
Week 27 - Freelance: $57820
Week 28 - Freelance: $57920
Week 29 - Freelance: $58020
Week 30 - Freelance: $58120
Week 31 - Freelance: $58220
Week 32 - Freelance: $58320
Week 33 - Freelance: $58420
Week 34 - Freelance: $58520
Week 35 - Freelance: $58620
Week 36 - Freelance: $58720
Week 37 - Freelance: $58820
Week 38 - Freelance: $58920
Week 39 - Freelance: $59020
Week 40 - Freelance: $59120
Week 41 - Freelance: $59220
Week 42 - Freelance: $59320
Week 43 - Freelance: $59420
Week 44 - Freelance: $59520
Week 45 - Freelance: $59620
Week 46 - Freelance: $59720
Week 47 - Freelance: $59820
Week 48 - Freelance: $59920
Week 49 - Freelance: $60020
Week 50 - Freelance: $60120
Week 51 - Freelance: $60220
Week 52 - Freelance: $60320
Final bankBalance: $60320

```



