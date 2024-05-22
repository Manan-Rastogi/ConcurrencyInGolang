package main

import (
	"fmt"
	"sync"
)

// 

// How much money will one make in next year..

type Income struct{
	Source string
	Amount int
}

func main(){
	// variable for bank balance
	var bankBalance int
	var wg sync.WaitGroup
	var mux sync.Mutex
	
	// print out starting values
	fmt.Printf("Initial Balance : %d.00\n", bankBalance)


	// define weekly revenue
	incomes := []Income{
		{Source: "JOB", Amount:12000},
		{Source: "Youtube", Amount:1000},
		{Source: "Investments", Amount:-100},
		{Source: "Freelancing", Amount:2500},

	}

	// loop through 52 weeks and print how much is made. Keep a running total
	for i, income := range(incomes){
		wg.Add(1)
		go func (i int, income Income)  {
			defer wg.Done()
			mux.Lock()
			for week:=1; week<=52;week++{
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				fmt.Printf("On week %d, you earned Rs %v.00 from %s.\n", week, income.Amount, income.Source)
			}
			mux.Unlock() 
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final Balance : Rs %d.00 \n", bankBalance)
}