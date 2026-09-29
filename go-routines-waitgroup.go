package main

import (
	"fmt"
	"sync"
)

func sayHello(name string) {
	fmt.Println("Hello:", name)
}

func main() {
	var wg sync.WaitGroup // Crea el grupo (nos da la sincronizacion)
						  // Pero WaitGroup no protege datos compartidos

	wg.Add(3) // le dice que se agregaran 3 

	go func() { // funcion anonima
		defer wg.Done() // defer: Se ejecuta al final de la funcion

		sayHello("Matias")
	}()

	go func() { // go es para crear una goroutine
		defer wg.Done()

		sayHello("Pedro")
	}()

	go func() {
		defer wg.Done()

		sayHello("Carla")
	}()

	wg.Wait()

	fmt.Println("Todos terminaron")
}

// Cuando se trabaja con concurrencia tenemos dos problemas
// Concurrencia -> Varias tareas progresan independientemente
// Sincronizacion -> Necesitamos coordinar esas tareas
