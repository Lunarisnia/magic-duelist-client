package main

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/entities"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/renderer"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
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
	playerPawn := entities.NewPawn(mtypes.Vector2i{X: 0, Y: 0})
	// TODO: this should immediately be updated from the server for the correct position
	opponentPawn := entities.NewPawn(mtypes.Vector2i{X: 59, Y: 0})
	// TODO: The world state should be obtained from the server
	world := world.NewWorld(playerPawn, opponentPawn)

	gameEngine := engine.NewGameEngine(s, gameRenderer, world)
	err = gameEngine.Start(ctx)
	if err != nil {
		panic(err)
	}
}
