package main

import (
	"fmt"
)

func inform() {
	fmt.Println("Hello from the meat store!")
}

func main() {
	fmt.Println("Hi!")
	go inform()
	fmt.Println("Bye!")
}

