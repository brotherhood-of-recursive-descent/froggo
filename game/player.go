package game

import (
	"froggo/lib"
)

type Player struct {
	lib.BaseEntity

	Position lib.Vec2
}

func NewPlayer(inputManager lib.InputManager) *Player {

	p := Player{}

	movement := PlayerMovement{
		Position:     &p.Position,
		InputManager: inputManager,
		Active:       true,
	}
	rectangle := NewPlayerRectangle(&p)

	p.AddComponent(&movement, rectangle)
	return &p
}
