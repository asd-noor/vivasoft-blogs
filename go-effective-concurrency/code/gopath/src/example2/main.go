package main

import (
	"fmt"
	"sync"
)

func inform(w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Hello from the meat store!")
}

func main() {
	fmt.Println("Hi!")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go inform(wg)

	wg.Wait()
	fmt.Println("Bye!")
}
