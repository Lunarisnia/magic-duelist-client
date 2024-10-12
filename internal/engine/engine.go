package engine

import (
	"context"
	"os"
	"time"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/renderer"
	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
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

func NewGameEngine(screen tcell.Screen, renderer renderer.Renderer, world world.World) GameEngine {
	return &GameEngineImpl{
		tick:     0,
		renderer: renderer,
		screen:   screen,
		world:    world,
	}
}

var directions = map[rune]mtypes.Vector2i{
	'a': mtypes.Vector2Left(),
	'd': mtypes.Vector2Right(),
	'w': mtypes.Vector2Up(),
	's': mtypes.Vector2Down(),
}

func (g *GameEngineImpl) Start(ctx context.Context) error {
	for {
		// Render
		g.screen.Clear()
		snapshot := g.world.GetSnapshot(ctx)
		g.renderer.Render(ctx, snapshot)
		g.screen.Show()

		// Ask for user input
		input := func() {
			for {
				ev := g.screen.PollEvent()

				switch ev := ev.(type) {
				case *tcell.EventResize:
					g.screen.Sync()
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyCtrlC {
						os.Exit(0)
					}
					if direction, exist := directions[ev.Rune()]; exist {
						g.world.MovePlayer(ctx, direction)
					}
					// NOTE: Might not be fixable since tcell don't report keyup
					// FIXME: Player can't shoot and move
					if ev.Rune() == ' ' {
						g.world.PlayerShooting(ctx)
					}
				}
			}
		}
		go input()

		g.world.MoveBullets(ctx)
		g.world.DestroyBullets(ctx)

		// Calculate result from input || send user input to server
		g.tick++
		time.Sleep(20 * time.Millisecond)
	}
}
