package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func ListenHTTP(serverState *ServerState) error {
	http.HandleFunc("/start-match", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Match Started")
	})
	// FIXME: This does not work
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Server should start a timer if the player isn't connected to the UDP server for a time then delete their ID
		if len(serverState.PlayersID) >= 2 {
			fmt.Fprint(w, "Room is full.")
			return
		}
		playerID := uuid.NewString()
		serverState.PlayersID = append(serverState.PlayersID, playerID)
		fmt.Fprint(w, playerID)
		fmt.Printf("Player has entered the room (%s)\n", playerID)
	})
	fmt.Println("HTTP Listening on 7000")
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		return err
	}

	return nil
}
