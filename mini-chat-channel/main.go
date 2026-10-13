package main

import (
	"fmt"
)

func main() {
	lobby := NewLobby()

	err := lobby.AddPlayer("Matias")

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = lobby.AddPlayer("Pedro")

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = lobby.AddPlayer("Carla")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Jugadores conectados:")

	go lobby.Run() // Empieza a escuchar mensajes

	err = lobby.SendMessage("Matias", "Hola")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = lobby.SendMessage("Pedro", "Buenas")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = lobby.SendMessage("Carla", "Como estan?")
	if err != nil {
		fmt.Println("Error:", err)
	}

}
