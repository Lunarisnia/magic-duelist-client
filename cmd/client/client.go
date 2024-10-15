package main

import (
	"context"
	"net"

	"github.com/Lunarisnia/magic-duelist-client/internal/client"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/entities"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/renderer"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
	"github.com/Lunarisnia/magic-duelist-client/internal/magicp"
	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
	"github.com/gdamore/tcell/v2"
)

// NOTE: Player should always be on the left from their point of view like Tekken, which mean I need to translate those movement properly before sending it to the client
// NOTE: Player should teleport to the other side if they goes out of bound vertically
// TODO: NEXT: Do the server and its protocol

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = s.Init()
	if err != nil {
		panic(err)
	}
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(style)
	s.Clear()
	// NOTE: if the panic doesnt tell us anything look at this https://github.com/gdamore/tcell/blob/main/TUTORIAL.md and search for "maybePanic"
	defer s.Fini()

	ctx := context.Background()
	gameRenderer := renderer.NewRenderer(s)

	// TODO: starting position should be coming from the server
	playerID, err := client.RegisterPlayer(ctx)
	if err != nil {
		panic(err)
	}
	playerPawn := entities.NewPawn(playerID, mtypes.Vector2i{X: 0, Y: 0})
	// TODO: this should immediately be updated from the server for the correct position
	opponentPawn := entities.NewPawn("", mtypes.Vector2i{X: 59, Y: 0})
	// TODO: The world state should be obtained from the server
	world := world.NewWorld(playerPawn, opponentPawn)

	udpAddr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 6900,
	}
	go magicp.Listen(udpAddr, func(snapshot *magicp.SnapshotProtocol) {
		// TODO: Translate player position to the player pointer directly
		newPlayerPosition := mtypes.Vector2i{
			X: snapshot.P1Position.X,
			Y: snapshot.P1Position.Y,
		}
		playerPawn.SetPosition(newPlayerPosition)

		// TODO: Translate opponent position to the opponent pointer directly (But translate the position first, remember opponent is always on the right)
		// TODO: Translate the bullet position from the server, hence I think each bullet should have its own ID and keyed to a map
	})

	gameEngine := engine.NewGameEngine(s, gameRenderer, world)
	err = gameEngine.Start(ctx)
	if err != nil {
		panic(err)
	}
}
