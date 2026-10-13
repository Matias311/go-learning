package main

import (
	"errors"
	"fmt"
)

type Lobby struct {
	Players  map[string]*Player
	Messages []Message // declaracion de slice
	Incoming chan Message
}

func NewLobby() *Lobby {
	return &Lobby{
		Players:  map[string]*Player{},
		Messages: []Message{}, // inicializacion de slice (naturalemente no hace falta pero asi es mas claro)
		Incoming: make(chan Message),
	}
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

// Se protege porque muchos pueden modificar el slice a la vez
func (lobby *Lobby) SendMessage(sender, content string) error {
	_, err := lobby.GetPlayer(sender)
	if err != nil {
		return err
	}

	message := Message{
		Sender:  sender,
		Content: content,
	}

	lobby.Incoming <- message

	return nil
}

func (lobby *Lobby) PlayerCount() int {
	return len(lobby.Players) // devuelve la cantidad de elementos
}

// Aqui se projete proque mientras alguien agrega un mensaje aun asi se puede leer
func (lobby *Lobby) MessageCount() int {
	return len(lobby.Messages)
}

// Aqui se protege porque se puede retornar la lista mientras se esta modificando pero no se puede proteger con
// RLock porque aunque se retorne despues de ser leido o modificado, no estamos retornando una copia, entonces
// estamos retornando el mismo slice que esta siendo moficado
func (lobby *Lobby) ListMessages() []Message {

	messages := make([]Message, len(lobby.Messages)) // crea un slice con el espacio necesario tomando
	// en cuenta el largo del slice de mensajes original
	copy(messages, lobby.Messages)

	return messages
}

func (Lobby *Lobby) Run() {
	for message := range Lobby.Incoming {
		Lobby.Messages = append(Lobby.Messages, message)
		fmt.Println("Procesando:", message.Sender, message.Content)
	}
}
