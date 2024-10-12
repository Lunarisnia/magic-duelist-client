package world

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/entities"
	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
)

const (
	Height = 15
	Width  = 60
)

type Arena [][]int

type Snapshot struct {
	Arena        Arena
	PlayerPawn   entities.Pawn
	OpponentPawn entities.Pawn
	BulletList   *BulletList
}

type BulletList struct {
	Bullet entities.Bullet
	Next   *BulletList
	Prev   *BulletList
}

type World interface {
	GetSnapshot(ctx context.Context) Snapshot
	MovePlayer(ctx context.Context, direction mtypes.Vector2i)
	PlayerShooting(ctx context.Context)
	OpponentShooting(ctx context.Context)
	MoveOpponent(ctx context.Context, direction mtypes.Vector2i)
	MoveBullets(ctx context.Context)
	DestroyBullets(ctx context.Context)
}

type WorldImpl struct {
	arena        Arena
	playerPawn   entities.Pawn
	opponentPawn entities.Pawn
	bulletList   *BulletList
	head         *BulletList
}

func NewWorld(playerPawn, enemyPawn entities.Pawn) World {
	arena := make(Arena, Height)
	for i := range arena {
		arena[i] = make([]int, Width)
	}
	head := BulletList{
		Bullet: nil,
		Next:   nil,
		Prev:   nil,
	}
	tail := BulletList{
		Bullet: nil,
		Prev:   nil,
		Next:   nil,
	}
	head.Next = &tail
	tail.Prev = &head
	return &WorldImpl{
		arena:        arena,
		playerPawn:   playerPawn,
		opponentPawn: enemyPawn,
		bulletList:   &head,
		head:         &head,
	}
}

func (w *WorldImpl) GetSnapshot(ctx context.Context) Snapshot {
	snapshot := Snapshot{
		Arena:        w.arena,
		PlayerPawn:   w.playerPawn,
		OpponentPawn: w.opponentPawn,
	}
	if w.bulletList != nil {
		snapshot.BulletList = w.bulletList
	}
	return snapshot
}

func (w *WorldImpl) MoveBullets(ctx context.Context) {
	// TODO: Should always inform the server where the bullet is and if its destroyed or hit something
	var move func(node *BulletList)
	move = func(node *BulletList) {
		if node == nil || node.Bullet == nil {
			return
		}

		node.Bullet.Move()
		move(node.Next)
	}

	move(w.bulletList.Next)
}

func (w *WorldImpl) MovePlayer(ctx context.Context, direction mtypes.Vector2i) {
	// TODO: should check to the server if its legal to move this way eg: not out of bound, if not then do nothing
	w.playerPawn.Move(ctx, direction)
}

func (w *WorldImpl) PlayerShooting(ctx context.Context) {
	bullet := w.playerPawn.Shoot(ctx, false)
	w.appendNode(bullet)
}

func (w *WorldImpl) OpponentShooting(ctx context.Context) {
	bullet := w.opponentPawn.Shoot(ctx, true)
	w.appendNode(bullet)
}

// MoveOpponent this will get called after the server had sent the last position of our opponent
func (w *WorldImpl) MoveOpponent(ctx context.Context, direction mtypes.Vector2i) {
	w.opponentPawn.Move(ctx, direction)
}

func (w *WorldImpl) appendNode(bullet entities.Bullet) {
	newNode := BulletList{
		Bullet: bullet,
		Next:   nil,
		Prev:   nil,
	}

	severedHead := w.head
	body := w.head.Next

	severedHead.Next = &newNode
	body.Prev = &newNode

	newNode.Next = body
	newNode.Prev = severedHead
}

func (w *WorldImpl) DestroyBullets(ctx context.Context) {
	var destroy func(node *BulletList)
	destroy = func(node *BulletList) {
		if node == nil || node.Bullet == nil {
			return
		}

		bulletPosition := node.Bullet.GetPosition()
		if bulletPosition.X >= Width || bulletPosition.X < 0 {
			severedHead := node.Prev
			severedBody := node.Next
			severedBody.Prev = severedHead
			severedHead.Next = severedBody
			return
		}
		destroy(node.Next)
	}
	if w.bulletList.Next != nil && w.bulletList.Next.Bullet != nil {
		destroy(w.bulletList.Next)
	}
}
