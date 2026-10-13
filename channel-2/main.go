package main

import "fmt"

type Message struct {
	Sender  string
	Content string
}

func main() {
	messages := make(chan Message)

	go func() {
		messages <- Message{
			Sender:  "Matias",
			Content: "Hola",
		}

		messages <- Message{
			Sender:  "Pedro",
			Content: "Bien y tu?",
		}

		messages <- Message{
			Sender:  "Carla",
			Content: "Como estan?",
		}

		close(messages)
	}()

	for message := range messages {
		fmt.Printf("%s: %s\n", message.Sender, message.Content)
	}

}
