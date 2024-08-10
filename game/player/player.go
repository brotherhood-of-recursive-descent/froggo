package player

import (
	"froggo/game"
	"froggo/lib"
)

type Player struct {
	game.Entity

	Position lib.Vec2
}

func NewPlayer() *Player {

	p := Player{}

	movement := PlayerMovement{}
	p.AddComponent(&movement)

	rectangle := PlayerRectangle{}
	p.AddComponent(&rectangle)

	return &p
}
