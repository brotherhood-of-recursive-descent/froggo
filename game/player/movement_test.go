package player_test

import (
	"froggo/game/player"
	"froggo/lib"
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
	p := player.NewPlayer()
	movement := player.PlayerMovement{
		Player:       p,
		InputManager: &fakeInput,
	}
	p.AddComponent(&movement)

	// trigger move left
	p.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = true
	fakeInput.MoveRightActive = false
	movement.Update()

	if p.Position.X != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, p.Position.X)
	}

	// trigger move right
	p.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = true
	movement.Update()

	if p.Position.X != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, p.Position.X)
	}

	// trigger move up
	p.Position = lib.Vec2{}
	fakeInput.MoveUpActive = true
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if p.Position.Y != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, p.Position.Y)
	}

	// trigger move down
	p.Position = lib.Vec2{}
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = true
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if p.Position.Y != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, p.Position.Y)
	}
}
