package player

import (
	"froggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

const MovementStep = 256

// PlayerMovement moves the player left, right, up or down
type PlayerMovement struct {
	InputManager lib.InputManager
	Player       *Player
	Active       bool
}

func (m *PlayerMovement) IsActive() bool {
	return m.Active
}

func (m *PlayerMovement) Update() error {

	if m.InputManager.MoveLeft() {
		m.Player.Position.X -= MovementStep
	}

	if m.InputManager.MoveRight() {
		m.Player.Position.X += MovementStep
	}

	if m.InputManager.MoveUp() {
		m.Player.Position.Y += MovementStep
	}

	if m.InputManager.MoveDown() {
		m.Player.Position.Y -= MovementStep
	}

	return nil
}

func (m *PlayerMovement) Draw(_ *ebiten.Image) { /*nothing to do*/ }
