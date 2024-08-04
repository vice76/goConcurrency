package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello world"

	wg.Add(2)
	go updateMessage("Hello , Universe !")
	go updateMessage("Hello , cosmos !")
	wg.Wait()

	fmt.Println(msg)

	//here , it is a problem with concurrency , as we don't have any idea
	//when which routine is going to finish , here line no 21 got finished
	//earlier than line no 20 , it also might be we get hello cosmos at sometime
	//how to check whether we are having race condition
	// go run -race .

}
