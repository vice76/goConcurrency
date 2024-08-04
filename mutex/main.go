package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello world"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello , Universe !", &mutex)
	go updateMessage("Hello , cosmos !", &mutex)
	wg.Wait()

	fmt.Println(msg)

	//here , it is a problem with concurrency , as we don't have any idea
	//when which routine is going to finish , here line no 21 got finished
	//earlier than line no 20 , it also might be we get hello cosmos at sometime
	//how to check whether we are having race condition
	// go run -race .

	// solution to this problm is using mutex to lock and unlock resource
	// its a thread safe operation , there will be no race condition

}
