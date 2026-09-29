package main 

import (
	"fmt"
	"sync"
)

func sayHello(name string) {
	fmt.Println("Hola", name)
}

func main() {

	go sayHello("Matias") // go significa aproximadamente: ejecuta esta funcion concurrentemente en una goroutine
	go sayHello("Pedro")
	go sayHello("Carla")
	// el problema actual es que termina y no se muestra nada porque le dice a go: empieza esta tarea y yo continuo
	// Cuando main termina, termina el programa completo. No espera automaticamente a las otras goroutine

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		sayHello("Matias")
	}()
	go func() {
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
