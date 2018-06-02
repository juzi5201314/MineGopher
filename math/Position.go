package math

type Position struct {
	X, Z int32
	Y    uint32
}

func NewPosition(x int32, y uint32, z int32) Position {
	return Position{X: x, Y: y, Z: z}
}
