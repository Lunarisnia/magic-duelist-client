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
	style  tcell.Style
}

func NewRenderer(screen tcell.Screen) Renderer {
	return &RendererImpl{
		screen: screen,
		style:  tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple),
	}
}

func (r *RendererImpl) Render(ctx context.Context, snapshot world.Snapshot) error {
	// TODO: render the world in basic ascii
	// TODO: Draw empty cell with blank space
	// TODO: Draw the players with a '>' and '<' respectively
	playerPosition := snapshot.PlayerPawn.GetPosition()
	opponentPosition := snapshot.OpponentPawn.GetPosition()
	for y, cols := range snapshot.Arena {
		for x := range cols {
			r.screen.SetContent(x, y, ' ', nil, r.style)
		}
	}
	r.drawPawn(playerPosition.X, playerPosition.Y, false)
	r.drawPawn(opponentPosition.X, opponentPosition.Y, true)
	r.drawBullets(snapshot.BulletList.Next)
	return nil
}

func (r *RendererImpl) drawPawn(x int, y int, isOpponent bool) {
	sprite := '>'
	if isOpponent {
		sprite = '<'
	}
	r.screen.SetContent(x, y, sprite, nil, r.style)
}

func (r *RendererImpl) drawBullets(bulletList *world.BulletList) {
	if bulletList == nil || bulletList.Bullet == nil {
		return
	}

	sprite := 'O'
	bulletPosition := bulletList.Bullet.GetPosition()
	r.screen.SetContent(bulletPosition.X, bulletPosition.Y, sprite, nil, r.style)
	r.drawBullets(bulletList.Next)
}
