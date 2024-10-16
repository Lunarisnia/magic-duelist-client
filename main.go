package main

import (
	"context"
	"fmt"

	"github.com/Lunarisnia/magic-duelist-client/internal/client"
)

// TODO: Do the start match API on the server, where you can assign each player to an ID
// TODO: Have their movement translated properly and is visibly moving well enough on both player

func main() {
	playerID, err := client.RegisterPlayer(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("PlayerID: ", playerID)
}

// NOTE: You left this project after encountering the problem of system design spefically how to check if both player has connected to the game. I think you need to take a step back and do something else before going back to this one.
