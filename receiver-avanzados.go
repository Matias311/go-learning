package main

import (
	"errors"
	"fmt"
)

type Player struct {
	Name      string
	Score     int
	Connected bool
}

type Lobby struct {
	Players map[string]*Player
}

func (player *Player) AddScore() {
	player.Score += 10
}

func NewLobby() *Lobby {
	return &Lobby{
		Players: map[string]*Player{},
	}
}

func (player *Player) Disconnect() {
	player.Connected = false
}

func (player *Player) Connect() {
	player.Connected = true
}

func (lobby *Lobby) AddPlayer(name string) error {
	_, exists := lobby.Players[name]

	if exists {
		return errors.New("El jugador ya se encuentra creado")
	}

	player := Player{
		Name:      name,
		Score:     0,
		Connected: true,
	}

	lobby.Players[name] = &player
	return nil
}

// reciver que retorna un objeto (retorno multiple)
func (lobby *Lobby) GetPlayer(name string) (*Player, error) {
	player, exists := lobby.Players[name]

	if !exists {
		return nil, errors.New("No se encontro al jugador")
	}

	return player, nil
}

// funcion delete nos ayuda a eliminar elementos del map
func (lobby *Lobby) RemovePlayer(name string) error {
	_, exists := lobby.Players[name]

	if !exists {
		return errors.New("No se encontro al jugador")
	}

	delete(lobby.Players, name)
	return nil
}

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

	for _, player := range lobby.Players {
		fmt.Println(player.Name, "- Score:", player.Score)

	}

	//TODO: Ejercicio: 1. Buscar a Matias 2. Eliminar a Pedro  3. Volver a buscar a Pedro
	player, err := lobby.GetPlayer("Matias")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Se encontro a:", player.Name)
	}

	err = lobby.RemovePlayer("Pedro")

	if err != nil {
		fmt.Println("Error:", err)
	}

	player, err = lobby.GetPlayer("Pedro")
	if err != nil {
		fmt.Println("Error:", err)
	}

}
