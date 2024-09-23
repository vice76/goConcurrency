package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
	//both forks will be mutex that is lock will be there
	//on both of them as per the problem
}

//list of philosopher

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define som variables
var hunger = 3 //how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	//printout a welcome message
	fmt.Println("Dining philosopher problem ")
	fmt.Println("Table is empty")

	//start the meal
	dine()

	//print out finished message
	fmt.Println("Table is empty.")
}

func dine() {
	//this waitgroup is for one philosopher had his meal
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	//we will only start after the all are seated
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is a map of all 5 forks
	//her we have lock the forks so for this
	// we will be using sync.mutex package
	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		//fire off the go routine for current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
	//pause all the routines as all are done
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
}
