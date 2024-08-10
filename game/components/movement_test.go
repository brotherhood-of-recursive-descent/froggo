package components_test

import (
	"froggo/game/components"
	"froggo/game/lib"
	"froggo/game/objects"
	"testing"
)

type FakeInputManager struct {
	MoveUpActive    bool
	MoveDownActive  bool
	MoveRightActive bool
	MoveLeftActive  bool
}

func (f *FakeInputManager) MoveUp() bool    { return f.MoveUpActive }
func (f *FakeInputManager) MoveDown() bool  { return f.MoveDownActive }
func (f *FakeInputManager) MoveLeft() bool  { return f.MoveLeftActive }
func (f *FakeInputManager) MoveRight() bool { return f.MoveRightActive }

func TestMovement(t *testing.T) {

	// setup
	fakeInput := FakeInputManager{}
	player := objects.NewPlayer()
	movement := components.PlayerMovement{
		Player:       player,
		InputManager: &fakeInput,
	}
	player.AddComponent(&movement)

	// trigger move left
	player.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = true
	fakeInput.MoveRightActive = false
	movement.Update()

	if player.Position.X != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, player.Position.X)
	}

	// trigger move right
	player.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = true
	movement.Update()

	if player.Position.X != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, player.Position.X)
	}

	// trigger move up
	player.Position = lib.Vec2{}
	fakeInput.MoveUpActive = true
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if player.Position.Y != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, player.Position.Y)
	}

	// trigger move down
	player.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = true
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if player.Position.Y != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, player.Position.Y)
	}
}
