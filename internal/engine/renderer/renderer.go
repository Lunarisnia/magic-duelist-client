package renderer

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
	"github.com/gdamore/tcell/v2"
)

type Renderer interface {
	Render(ctx context.Context, snapshot world.Snapshot) error
}

type RendererImpl struct {
	screen tcell.Screen
}

func NewRenderer(screen tcell.Screen) Renderer {
	return &RendererImpl{
		screen: screen,
	}
}

func (r *RendererImpl) Render(ctx context.Context, snapshot world.Snapshot) error {
	// TODO: render the world in basic ascii
	// TODO: Draw empty cell with blank space
	// TODO: Draw the player with a 'O'
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)
	for y, cols := range snapshot.Arena {
		for x := range cols {
			r.screen.SetContent(x, y, '0', nil, style)
		}
	}
	return nil
}
