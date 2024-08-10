package player

import "github.com/hajimehoshi/ebiten/v2"

// #2 - add a new type satisfying the component interface
// then implement the rendering in the Draw(ebiten.Image) function

type PlayerRectangle struct{}

func (p *PlayerRectangle) IsActive() bool {
	panic("not implemented") // TODO: Implement
}

func (p *PlayerRectangle) Update() error {
	panic("not implemented") // TODO: Implement
}

func (p *PlayerRectangle) Draw(_ *ebiten.Image) {
	panic("not implemented") // TODO: Implement
}
