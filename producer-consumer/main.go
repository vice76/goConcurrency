package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
	//channnels of channel
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved prder #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++

		fmt.Printf("Makaing pizza #%d.It will take %d seconds...\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)

		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** Cook quit for making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	} else {
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
		}
	}
}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0

	//run forever or until we recieve a quit notification

	//try to make pizzas

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				//close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	//In our pc problem , producer is a pizza manufacturer
	//and consumer is a customer

	//seed the random number generator for generating pizza
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("the pizzeria is open for business!")
	color.Cyan("------------------")

	// craete a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for order", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("the customer is really mad !")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	//print out the ending message

	color.Cyan("------------------------")
	color.Cyan("Done for the day")

	color.Cyan("We made #%d pizzas , but failed to make #%d , with attempts #%d in total.", pizzaMade, pizzaFailed, total)

	switch {
	case pizzaFailed > 9:
		color.Red("It was an awful day")
	case pizzaFailed >= 6:
		color.Red("It was not a very good day")
	case pizzaFailed >= 4:
		color.Yellow("It was an okay day")
	case pizzaFailed >= 2:
		color.Yellow("It was pretty good day")
	default:
		color.Green("best day")

	}

}
