package entities

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
)

type Pawn interface {
	Move(ctx context.Context, direction mtypes.Vector2i)
	GetPosition() mtypes.Vector2i
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
