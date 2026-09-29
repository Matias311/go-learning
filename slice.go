package main

import (
	"errors"
	"fmt"
)

type Message struct {
	Sender  string
	Content string
}

type Player struct {
	Name      string
	Score     int
	Connected bool
}

type Lobby struct {
	Players  map[string]*Player
	Messages []Message // declaracion de slice
}

func (player *Player) AddScore() {
	player.Score += 10
}

func NewLobby() *Lobby {
	return &Lobby{
		Players:  map[string]*Player{},
		Messages: []Message{}, // inicializacion de slice (naturalemente no hace falta pero asi es mas claro)
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

func (lobby *Lobby) SendMessage(sender, content string) error {
	_, err := lobby.GetPlayer(sender)
	if err != nil {
		return err
	}

	message := Message{
		Sender:  sender,
		Content: content,
	}

	lobby.Messages = append(lobby.Messages, message) // reescribimos Mensajes dentro de loby porque append devuelve el slice actualizado, por eso debemos reasignarlo
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

	// TODO: Enviar mensajes
	lobby.SendMessage("Matias", "Hola a todos!")
	lobby.SendMessage("Carla", "Hola Matias!")
	lobby.SendMessage("Pedro", "Que tal?")

	for _, message := range lobby.Messages {
		fmt.Println(message.Sender, ":", message.Content)
	}

	err = lobby.SendMessage("Pepito", "Soy Hacker")

	if err != nil {
		fmt.Println("error:", err)
	}
}
