package world

import (
	"context"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/entities"
	"github.com/Lunarisnia/magic-duelist-client/internal/mtypes"
)

type Arena [][]int

type Snapshot struct {
	Arena        Arena
	PlayerPawn   entities.Pawn
	OpponentPawn entities.Pawn
}

type World interface {
	GetSnapshot(ctx context.Context) Snapshot
	MovePlayer(ctx context.Context, direction mtypes.Vector2i)
	MoveOpponent(ctx context.Context, direction mtypes.Vector2i)
}

type WorldImpl struct {
	arena        Arena
	playerPawn   entities.Pawn
	opponentPawn entities.Pawn
}

func NewWorld(playerPawn, enemyPawn entities.Pawn) World {
	arena := make(Arena, 15)
	for i := range arena {
		arena[i] = make([]int, 60)
	}
	return &WorldImpl{
		arena:        arena,
		playerPawn:   playerPawn,
		opponentPawn: enemyPawn,
	}
}

func (w *WorldImpl) GetSnapshot(ctx context.Context) Snapshot {
	return Snapshot{
		Arena:        w.arena,
		PlayerPawn:   w.playerPawn,
		OpponentPawn: w.opponentPawn,
	}
}

func (w *WorldImpl) MovePlayer(ctx context.Context, direction mtypes.Vector2i) {
	// TODO: should check to the server if its legal to move this way eg: not out of bound, if not then do nothing
	w.playerPawn.Move(ctx, direction)
}

// MoveOpponent this will get called after the server had sent the last position of our opponent
func (w *WorldImpl) MoveOpponent(ctx context.Context, direction mtypes.Vector2i) {
	w.opponentPawn.Move(ctx, direction)
}
