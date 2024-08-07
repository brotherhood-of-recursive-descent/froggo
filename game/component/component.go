package component

import "github.com/hajimehoshi/ebiten/v2"

// Component is the basic building block of an entity. An entity can hold multiple components. Components can be reused across different entities.
type Component interface {
	IsActive() bool
	Update() error
	Draw(*ebiten.Image)
}
