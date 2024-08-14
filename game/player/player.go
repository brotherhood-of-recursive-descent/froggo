package player

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
		Player:       &p,
		InputManager: inputManager,
	}
	rectangle := NewPlayerRectangle(&p)

	p.AddComponent(&movement, rectangle)
	return &p
}
