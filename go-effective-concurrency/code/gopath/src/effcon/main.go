package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Hello from the other side")
}

func main() {
	go greet()
	fmt.Println("Hello")
}
