package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
	ProductCount int
}

func addToInventory(w *sync.WaitGroup, inventory *Inventory) {
	defer w.Done()
	inventory.ProductCount++
}

func main() {
	fmt.Println("Collecting the products...")
	wg := &sync.WaitGroup{}
	inventory := Inventory{
		ProductCount: 0,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go addToInventory(wg, &inventory)
	}

	wg.Wait()
	fmt.Printf("There are %d products in inventory\n\n", inventory.ProductCount)
}
