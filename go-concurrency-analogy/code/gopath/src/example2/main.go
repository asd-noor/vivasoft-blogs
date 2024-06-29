package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchTickets(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Got your tickets. Coming...")
	time.Sleep(3 * time.Second)
	fmt.Println("Reached airport")
}

func main() {
	fmt.Println("I left my tickets at home, bring them to me!")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fetchTickets(wg)

	wg.Wait()
	fmt.Println("Thanks! Adios, amigo.")
}
