package main 

type Player struct {
	Name      string
	Score     int
	Connected bool
}

func (player *Player) AddScore() {
	player.Score += 10
}

func (player *Player) Disconnect() {
	player.Connected = false
}

func (player *Player) Connect() {
	player.Connected = true
}
