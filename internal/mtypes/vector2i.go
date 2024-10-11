package mtypes

type Vector2i struct {
	X int
	Y int
}

func Vector2Right() Vector2i {
	return Vector2i{
		X: 1,
		Y: 0,
	}
}

func Vector2Left() Vector2i {
	return Vector2i{
		X: -1,
		Y: 0,
	}
}

func Vector2Up() Vector2i {
	return Vector2i{
		X: 0,
		Y: -1,
	}
}

func Vector2Down() Vector2i {
	return Vector2i{
		X: 0,
		Y: 1,
	}
}

func (v *Vector2i) Add(b Vector2i) {
	v.X += b.X
	v.Y += b.Y
}
