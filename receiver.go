package main

import (
	"errors"
	"fmt"
)

// Jugadores conectados:
// Matias - Score: 0
// Pedro - Score: 20
// Carla - Score: 40
// Actualizando puntaje de Matias...
// Matias - Score: 10

type Player struct {
	Name      string
	Score     int
	Connected bool
}

type Lobby struct {
	Players map[string]*Player
}

// Los receiver podemos persona que son metodos de una estructura
// AddScore es un metodo que pertenece a *Player
// Si no utilizamos *Player vamos a modificar una copia, por eso se le pasa la referencia en memoria
// a eso se le llama pointer receiver
func (player *Player) AddScore() {
	player.Score += 10
}

// Crea una nueva lobby (como buena practica se empieza con New).
// Retorna un puntero de lobby
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

// Agrega un jugador, si el jugador ya esta en la lobby lanza un error, si no lo esta se crea y se guarda retornando nil
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

	player, exists := lobby.Players["Matias"]
	if exists {
		fmt.Println("Actualizando puntaje de Matias...")
		player.AddScore()
		fmt.Println(player.Name, "- Score:", player.Score)

		fmt.Println("Estado:", player.Connected)
		player.Disconnect()
		fmt.Println("Estado:", player.Connected)
		player.Connect()
		fmt.Println("Estado:", player.Connected)
	}


	// Prueba, deberia de saltar el error
	err = lobby.AddPlayer("Pedro")

	if err != nil {
		fmt.Println("Error:", err)
	}
}
