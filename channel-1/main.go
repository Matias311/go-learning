package main

import "fmt"

func main() {
	numbers := make(chan int)

	go func() {
		numbers <- 42
	}()

	number := <-numbers

	fmt.Print("Numero recibido:", number)
}
