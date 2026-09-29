package main

import "errors"

type Lobby struct {
	Players  map[string]*Player
	Messages []Message // declaracion de slice
}

func NewLobby() *Lobby {
	return &Lobby{
		Players:  map[string]*Player{},
		Messages: []Message{}, // inicializacion de slice (naturalemente no hace falta pero asi es mas claro)
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

func (lobby *Lobby) PlayerCount() int{
	return len(lobby.Players) // devuelve la cantidad de elementos
}

func (lobby *Lobby) MessageCount() int {
	return len(lobby.Messages)
}

func (lobby *Lobby) ListMessages() []Message {
	return lobby.Messages
}
