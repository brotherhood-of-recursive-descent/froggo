package game

import (
	"froggo/lib"
)

type Player struct {
	lib.BaseEntity

	Position lib.Vec2
}

func NewPlayer(world World) *Player {

	p := Player{}

	p.Position = startPosition(world.ScreenDimension)
	movement := NewPlayerMovement(world, &p)
	rectangle := NewPlayerRectangle(&p)

	p.AddComponent(&movement, rectangle)
	return &p
}

func startPosition(screenDimension lib.Vec2) lib.Vec2 {
	return lib.Vec2{
		X: (screenDimension.X - RectangleWidthAndHeight) / 2,
		Y: screenDimension.Y - RectangleWidthAndHeight - 10,
	}
}
