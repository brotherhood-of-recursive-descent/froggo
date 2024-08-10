package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	KEY_UP    = ebiten.KeyW
	KEY_LEFT  = ebiten.KeyA
	KEY_DOWN  = ebiten.KeyS
	KEY_RIGHT = ebiten.KeyD
)

// InputManager wraps the input system behind concrete methods
type InputManager interface {
	MoveLeft() bool
	MoveRight() bool
	MoveUp() bool
	MoveDown() bool
}

// EbitenInputManager implements InputManager using the ebitenengine as a input source
type EbitenInputManager struct{}

func (i *EbitenInputManager) MoveLeft() bool {
	return ebiten.IsKeyPressed(KEY_LEFT)
}

func (i *EbitenInputManager) MoveRight() bool {
	return ebiten.IsKeyPressed(KEY_RIGHT)
}

func (i *EbitenInputManager) MoveDown() bool {
	return ebiten.IsKeyPressed(KEY_DOWN)
}

func (i *EbitenInputManager) MoveUp() bool {
	return ebiten.IsKeyPressed(KEY_UP)
}
