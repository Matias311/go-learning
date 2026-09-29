package main 
import "fmt"

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

	lobby.SendMessage("Matias", "Hola a todos!")
	lobby.SendMessage("Carla", "Hola Matias!")
	lobby.SendMessage("Pedro", "Que tal?")

	for _, message := range lobby.Messages {
		fmt.Println(message.Sender, ":", message.Content)
	}
	
	fmt.Println("Cantidad de jugadores:", lobby.PlayerCount())
	fmt.Println("Cantidad de mensajes enviados:", lobby.MessageCount())

	messages := lobby.ListMessages()
	for _, message := range lobby.Messages {
		fmt.Println(message.Sender, ":", message.Content)
	}

	fmt.Println("len:", len(messages))
	fmt.Println("cap:", cap(messages)) // cap indica elementos caben en el array interno actual
									   // antes de que append pueda necesitar reservar otro array
									   // es variable la cantidad que deja asignada

	messages[0].Content = "MENSAJE MODIFICADO"
	fmt.Println(lobby.Messages[0].Content)

}
