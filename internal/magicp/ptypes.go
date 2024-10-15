package magicp

import (
	"encoding/json"

	"github.com/Lunarisnia/magic-duelist-client/internal/engine/world"
)

type PositionProtocol struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type BulletProtocol struct {
	Position  PositionProtocol `json:"position"`
	Direction PositionProtocol `json:"direction"`
}

type SnapshotProtocol struct {
	Author     int              `json:"author"`
	P1Position PositionProtocol `json:"p1_position"`
	P2Position PositionProtocol `json:"p2_position"`
	Bullets    []BulletProtocol `json:"bullets"`
}

func Unmarshal(worldSnapshot world.Snapshot) ([]byte, error) {
	p1Position := worldSnapshot.PlayerPawn.GetPosition()
	p2Position := worldSnapshot.OpponentPawn.GetPosition()

	snapshotProtocol := SnapshotProtocol{
		Author: worldSnapshot.Author,
		P1Position: PositionProtocol{
			X: p1Position.X,
			Y: p1Position.Y,
		},
		P2Position: PositionProtocol{
			X: p2Position.X,
			Y: p2Position.Y,
		},
	}

	bullets := make([]BulletProtocol, 0)

	var walk func(b *world.BulletList)
	walk = func(b *world.BulletList) {
		if b == nil {
			return
		}
		bPosition := PositionProtocol{
			X: b.Bullet.GetPosition().X,
			Y: b.Bullet.GetPosition().Y,
		}
		bDirection := PositionProtocol{
			X: b.Bullet.GetDirection().X,
			Y: b.Bullet.GetDirection().Y,
		}
		bullet := BulletProtocol{
			Position:  bPosition,
			Direction: bDirection,
		}
		bullets = append(bullets, bullet)

		walk(b.Next)
	}
	walk(worldSnapshot.BulletList)

	b, err := json.Marshal(snapshotProtocol)
	if err != nil {
		return nil, err
	}

	return b, nil
}
