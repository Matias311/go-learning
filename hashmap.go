package main

import "fmt"

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

func addScore(player *Player) {
	player.Score += 10
}

func main() {
	matias := Player{
		Name:      "Matias",
		Score:     0,
		Connected: true,
	}

	pedro := Player{
		Name:      "Pedro",
		Score:     20,
		Connected: true,
	}

	carla := Player{
		Name:      "Carla",
		Score:     40,
		Connected: true,
	}

	players := map[string]*Player{}
	players["Matias"] = &matias
	players["Pedro"] = &pedro
	players["Carla"] = &carla

	fmt.Println("Jugadores conectados:")

	for _, player := range players {
		fmt.Println(player.Name, "- Score:", player.Score)

	}

	player, exists := players["Matias"]
	if exists {
		fmt.Println("Actualizando puntaje de Matias...")
		addScore(player) // aqui ya viene como puntero
		fmt.Println(player.Name, "- Score:", player.Score)
	}
}
