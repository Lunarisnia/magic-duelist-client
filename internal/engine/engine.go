package engine

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/renderer"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
	"github.com/gdamore/tcell/v2"
)

type GameEngineImpl struct {
	tick     uint64
	renderer renderer.Renderer
	screen   tcell.Screen
	world    world.World
}

type GameEngine interface {
	Start(ctx context.Context) error
}

// TODO: Need a renderer service, and players
func NewGameEngine(screen tcell.Screen, renderer renderer.Renderer, world world.World) GameEngine {
	return &GameEngineImpl{
		tick:     0,
		renderer: renderer,
		screen:   screen,
		world:    world,
	}
}

func (g *GameEngineImpl) Start(ctx context.Context) error {
	for {
		// Render
		g.screen.Clear()
		snapshot := g.world.GetSnapshot(ctx)
		g.renderer.Render(ctx, snapshot)
		g.screen.Show()

		// Ask for user input
		ev := g.screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return nil
			}
		}

		// Calculate result from input || send user input to server
		g.tick++
	}
}
