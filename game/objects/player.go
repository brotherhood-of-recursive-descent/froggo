package objects

import (
	"froggo/game"
	"froggo/game/lib"
)

type Player struct {
	game.Entity

	Position lib.Vec2
}

func NewPlayer() *Player {

	p := Player{}
	// #2 - here you can add your component
	// p.AddComponent(myComponent)

	return &p
}
