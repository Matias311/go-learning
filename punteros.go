package main

import "fmt"

// Queremos esto:
// Usuario: Matias
// Puntaje inicial: 0
// Nueva puntuacion: 10
// Nueva puntuacion: 20
// Nueva puntuacion: 30
// Nueva puntuacion: 40
// Nueva puntuacion: 50
// Tengo que utilizar name := "Matias"

type Player struct {
	Name      string
	Score     int
	Connected bool
}

// En go los argumentos se pasan por valor
// Es decir, la funcion recibe una copia del valor que le entregas
// En este caso nos queda super para utilizar un puntero!
// le queremos decir: te voy a decir donde esta player original. Modificalo directamente
// eso es un puntero. Un puntero es una varialbe que contine la direccion donde esta en memoria
func addScore(player *Player) { // Ese * dice que recibe un puntero de player
	player.Score += 10
}

func main() {
	player := Player{
		Name:      "Matias",
		Score:     0,
		Connected: true,
	}

	fmt.Println("Usuario: ", player.Name)
	fmt.Println("Puntaje inicial: ", player.Score)

	for player.Score < 50 {
		addScore(&player) // &player obtiene la direccion de meoria de player y esa direccion es la que recibe addScore
		fmt.Println("Nueva puntuacion: ", player.Score)
	}

	fmt.Println("Conectado: ", player.Connected)
}

// Piezas fundamentales de punteros
// Por ahora &player se asocia con: dame la direccion de player
// y *Player: este valor es un puntero hacia un Player
