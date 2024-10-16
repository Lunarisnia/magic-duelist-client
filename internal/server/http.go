package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func handleStartMatch(serverState *ServerState) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(serverState.PlayersID) != 2 {
			fmt.Fprint(w, "We are still waiting for 1 or more players.")
			return
		}

		fmt.Fprint(w, "200")
	}
}

func handleRegister(serverState *ServerState) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Server should start a timer if the player isn't connected to the UDP server for a time then delete their ID
		if len(serverState.PlayersID) >= 2 {
			fmt.Fprint(w, "Room is full.")
			return
		}
		playerID := uuid.NewString()
		serverState.PlayersID = append(serverState.PlayersID, playerID)
		fmt.Fprint(w, playerID)
		fmt.Printf("Player has entered the room (%s)\n", playerID)
	}
}

func ListenHTTP(serverState *ServerState) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/start-match", handleStartMatch(serverState))
	mux.HandleFunc("/register", handleRegister(serverState))

	fmt.Println("HTTP Listening on 7000")
	err := http.ListenAndServe(":7000", mux)
	if err != nil {
		return err
	}

	return nil
}
