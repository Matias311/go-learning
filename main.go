package main

import "fmt" // fmt es el paquete de standard library

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hola desde GO")

	// Estructura de las variables 
	// var + nombre_var + data_type + = + value
	// var name string = "Matias"
	// var age int = 23
	// var connected bool = true

	// Es lo mismo que arriba pero con inferencia de dato
	// esta forma de inferencia es fuertemente tipada
	name := "Matias" 
	age := 23
	connected := true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(connected)

	// llamado de la funcion
	// Estructura de funcion
	// func + nombre + ( variable + tipo de dato) + retorno
	fmt.Println(add(1, 2))


	// if + variable + statement
	age:= 23

	if age >= 18 {
		fmt.Println("Mayor de edad")
	} else {
		fmt.Println("Menor de edad")
	}


	// Bucles
	// GO no tiene un while separado, es un for
	for i:= 0; i < 5; i++ {
		fmt.Println("Hola desde un bucle for: " + i)
	}

	// While 
	counter := 0
	for counter < 5 {
		fmt.Println("Hola desde un for while en go: " + i) 
		counter++
	}

	// for infinito
	// for {}
}


