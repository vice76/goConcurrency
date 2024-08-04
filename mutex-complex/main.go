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

func main() {
	//variable for bank balance
	var bankBalance int

	var balance sync.Mutex

	//print out starting values
	fmt.Printf("Initial account balance :$%d.00", bankBalance)

	//define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gift", Amount: 10},
		{Source: "part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	//loop through 52 weeks and print out how much is made
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 0; week < 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d,you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	//print out final baalnce

	fmt.Printf("Final bank balance $%d.00", bankBalance)

}
