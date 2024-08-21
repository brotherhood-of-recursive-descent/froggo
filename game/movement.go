package game

import (
	"froggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

const MovementStep = 256

// PlayerMovement moves the player left, right, up or down
type PlayerMovement struct {
	InputManager lib.InputManager
	Position     *lib.Vec2
	Active       bool
}

func NewPlayerMovement(world World, p *Player) PlayerMovement {

	return PlayerMovement{
		Position:     &p.Position,
		InputManager: world.InputManager,
		Active:       true,
	}
}

func (m *PlayerMovement) IsActive() bool {
	return m.Active
}

func (m *PlayerMovement) Update() error {
	if m.InputManager.MoveLeft() {
		m.Position.X -= MovementStep
	}

	if m.InputManager.MoveRight() {
		m.Position.X += MovementStep
	}

	if m.InputManager.MoveUp() {
		m.Position.Y -= MovementStep
	}

	if m.InputManager.MoveDown() {
		m.Position.Y += MovementStep
	}

	return nil
}

func (m *PlayerMovement) Draw(_ *ebiten.Image) { /*nothing to do*/ }
