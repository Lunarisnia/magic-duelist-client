package entities

import "github.com/Lunarisnia/magic-duelist-client/internal/mtypes"

type Bullet interface {
	Move()
	GetPosition() mtypes.Vector2i
}

type BulletImpl struct {
	owner     Pawn
	position  mtypes.Vector2i
	direction mtypes.Vector2i
}

func NewBullet(owner Pawn, origin mtypes.Vector2i, direction mtypes.Vector2i) Bullet {
	bullet := BulletImpl{
		position:  origin,
		owner:     owner,
		direction: direction,
	}
	return &bullet
}

func (b *BulletImpl) Move() {
	b.position.Add(b.direction)
}

func (b *BulletImpl) GetPosition() mtypes.Vector2i {
	return b.position
}
