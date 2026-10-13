package main

import (
	"errors"
	"fmt"
)

type ListMessageRequest struct {
	Result chan []Message
}

type MessageCountRequest struct {
	Result chan int
}

type UnregisterRequest struct {
	Name   string
	Result chan error
}

type GetPlayerResult struct {
	Player *Player
	Err    error
}

type PlayerCountRequest struct {
	Result chan int
}

type Lobby struct {
	// Estado
	Players  map[string]*Player
	Messages []Message // declaracion de slice

	// Canales de comunicacion
	Incoming     chan Message
	Register     chan RegisterRequest
	Unregister   chan UnregisterRequest
	FindPlayer   chan GetPlayerRequest
	PlayerCount  chan PlayerCountRequest
	MessageCount chan MessageCountRequest
	ListMessages chan ListMessageRequest
}

type RegisterRequest struct {
	Player *Player
	Result chan error
}

type GetPlayerRequest struct {
	Name   string
	Result chan GetPlayerResult
}

func NewLobby() *Lobby {
	return &Lobby{
		Players:      map[string]*Player{},
		Messages:     []Message{}, // inicializacion de slice (naturalemente no hace falta pero asi es mas claro)
		Incoming:     make(chan Message),
		Register:     make(chan RegisterRequest),
		Unregister:   make(chan UnregisterRequest),
		FindPlayer:   make(chan GetPlayerRequest),
		PlayerCount:  make(chan PlayerCountRequest),
		MessageCount: make(chan MessageCountRequest),
		ListMessages: make(chan ListMessageRequest),
	}
}

func (lobby *Lobby) GetPlayer(name string) (*Player, error) {
	result := make(chan GetPlayerResult)

	lobby.FindPlayer <- GetPlayerRequest{
		Name:   name,
		Result: result,
	}

	response := <-result // esperemos a que nos llegue un GetPlayerResult del channel

	return response.Player, response.Err
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

	lobby.Incoming <- message

	return nil
}

func (lobby *Lobby) GetPlayerCount() int {
	result := make(chan int)

	lobby.PlayerCount <- PlayerCountRequest{
		Result: result,
	}

	return <-result
}

func (lobby *Lobby) GetMessageCount() int {
	result := make(chan int)

	lobby.MessageCount <- MessageCountRequest{
		Result: result,
	}

	return <-result
}

func (lobby *Lobby) GetListMessages() []Message {

	result := make(chan []Message)

	lobby.ListMessages <- ListMessageRequest{
		Result: result,
	}

	return <-result
}

func (lobby *Lobby) RegisterPlayer(name string) error {

	player := Player{
		Name:      name,
		Score:     0,
		Connected: true,
	}

	result := make(chan error)

	lobby.Register <- RegisterRequest{
		Player: &player,
		Result: result,
	}

	return <-result
}

func (lobby *Lobby) UnregisterPlayer(name string) error {
	result := make(chan error)

	lobby.Unregister <- UnregisterRequest{
		Name:   name,
		Result: result,
	}

	return <-result
}
func (lobby *Lobby) Run() {
	for {
		select {
		case request := <-lobby.Register:
			_, exists := lobby.Players[request.Player.Name]

			if exists {
				request.Result <- errors.New("el jugador ya existe")
				continue
			}

			lobby.Players[request.Player.Name] = request.Player
			fmt.Println("Registrando:", request.Player.Name)
			request.Result <- nil

		case request := <-lobby.Unregister:
			_, exists := lobby.Players[request.Name]
			if !exists {
				request.Result <- errors.New("el jugador no se ha encontrado")
				continue
			}

			delete(lobby.Players, request.Name)
			fmt.Println("Eliminando:", request.Name)
			request.Result <- nil

		case message := <-lobby.Incoming:
			lobby.Messages = append(lobby.Messages, message)
			fmt.Println("Messaje:", message.Sender, message.Content)

		case request := <-lobby.FindPlayer:
			player, exists := lobby.Players[request.Name]

			if !exists {
				request.Result <- GetPlayerResult{
					Player: nil,
					Err:    errors.New("jugador no encontrado"),
				}

				continue
			}

			request.Result <- GetPlayerResult{
				Player: player,
				Err:    nil,
			}
		case request := <-lobby.PlayerCount:
			request.Result <- len(lobby.Players)
		case request := <-lobby.MessageCount:
			request.Result <- len(lobby.Messages)

		case request := <-lobby.ListMessages:
			messages := make([]Message, len(lobby.Messages))
			copy(messages, lobby.Messages)
			request.Result <- messages
		}

	}
}
