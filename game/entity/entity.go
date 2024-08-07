package entity

import (
	"errors"
	"froggo/game/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	component.DefaultContainer
}

func (e *Entity) Update() error {

	for k, c := range e.AllComponents() {
		if c.IsActive() {
			if err := c.Update(); err != nil {
				return errors.Join(errors.New("failed calling Update() on: '"+k), err)
			}
		}
	}

	return nil
}

func (e *Entity) Draw(screen *ebiten.Image) {

	for _, c := range e.AllComponents() {
		if c.IsActive() {
			c.Draw(screen)
		}
	}
}
