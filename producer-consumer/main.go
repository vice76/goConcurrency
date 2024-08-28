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
		//try to make a pizza
		//decision
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

	//print out the ending message

}
