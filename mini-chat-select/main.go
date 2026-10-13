package main

import (
	"fmt"
)

func main() {
	lobby := NewLobby()
	go lobby.Run() // Empieza a escuchar mensajes

	err := lobby.RegisterPlayer("Matias")
	if err != nil {
		fmt.Println("No se pudo registrar")
	} else {
		fmt.Println("Se registro!")
	}

	err = lobby.RegisterPlayer("Pedro")
	if err != nil {
		fmt.Println("No se pudo registrar")
	} else {
		fmt.Println("Se registro!")
	}

	err = lobby.RegisterPlayer("Carla")
	if err != nil {
		fmt.Println("No se pudo registrar")
	} else {
		fmt.Println("Se registro!")
	}

	err = lobby.SendMessage("Matias", "Hola")
	if err != nil {
		fmt.Println("No se pudo enviar mensaje")
	}

	err = lobby.SendMessage("Pedro", "Buenas")
	if err != nil {
		fmt.Println("No se pudo enviar mensaje")
	}

	err = lobby.UnregisterPlayer("Pedro")
	if err != nil {
		fmt.Println("No se pude eliminar")
	}

	player, error := lobby.GetPlayer("Matias")
	if error != nil {
		fmt.Println("No se pudo obtener al jugador")
	} else {
		fmt.Println(player.Name)
	}

}
