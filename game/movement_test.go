package game_test

import (
	"fmt"
	"froggo/game"
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
	p := &lib.Vec2{}

	movement := game.PlayerMovement{
		Position:     p,
		InputManager: &fakeInput,
	}

	// trigger move left
	p = &lib.Vec2{} // reset
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = true
	fakeInput.MoveRightActive = false
	movement.Update()

	fmt.Printf("%v", p)

	if p.X != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, p.X)
	}

	// trigger move right
	p = &lib.Vec2{} // reset
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = true
	movement.Update()

	if p.X != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, p.X)
	}

	// trigger move up
	p = &lib.Vec2{} // reset
	fakeInput.MoveUpActive = true
	fakeInput.MoveDownActive = false
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if p.Y != -256 {
		t.Errorf("expected new position to be %v, got %v\n", -256, p.Y)
	}

	// trigger move down
	p = &lib.Vec2{} // reset
	fakeInput.MoveUpActive = false
	fakeInput.MoveDownActive = true
	fakeInput.MoveLeftActive = false
	fakeInput.MoveRightActive = false
	movement.Update()

	if p.Y != 256 {
		t.Errorf("expected new position to be %v, got %v\n", 256, p.Y)
	}
}
