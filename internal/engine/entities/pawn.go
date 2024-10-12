package entities

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
)

type Pawn interface {
	Move(ctx context.Context, direction mtypes.Vector2i)
	GetPosition() mtypes.Vector2i
	Shoot(ctx context.Context, isOpponent bool) Bullet
}

type PawnImpl struct {
	position mtypes.Vector2i
}

func NewPawn(origin mtypes.Vector2i) Pawn {
	return &PawnImpl{
		position: origin,
	}
}

func (p *PawnImpl) Move(ctx context.Context, direction mtypes.Vector2i) {
	p.position.Add(direction)
}

func (p *PawnImpl) GetPosition() mtypes.Vector2i {
	return p.position
}

func (p *PawnImpl) Shoot(ctx context.Context, isOpponent bool) Bullet {
	bulletOrigin := mtypes.Vector2i{}
	bulletOrigin.Add(p.position)
	bulletDirection := mtypes.Vector2Right()
	if isOpponent {
		bulletOrigin.X -= 1
		bulletDirection = mtypes.Vector2Left()
	} else {
		bulletOrigin.X += 1
	}
	bullet := NewBullet(p, bulletOrigin, bulletDirection)
	return bullet
}
