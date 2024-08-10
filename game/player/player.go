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

	// #2 - here you can add your component
	// rectancle := PlayerRectangle{}
	// p.Addcomponent(rectangle)

	return &p
}
