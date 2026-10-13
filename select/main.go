package main

import "fmt"

func main() {
	messages := make(chan string)
	numbers := make(chan int)

	go func() {
		messages <- "Hola"
	}()

	go func() {
		numbers <- 42
	}()

	for range 2 {
		select {
		case message := <-messages:
			fmt.Println("Mensaje recibido:", message)
		case number := <-numbers:
			fmt.Println("Numero recibido:", number)
		}

	}
}
