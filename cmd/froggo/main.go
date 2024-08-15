package main

import (
	"froggo/game"
	"froggo/lib"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// to be replaced by build script
var VERSION = "latest"

// FroggoGame contains all global available objects
type FroggoGame struct {
	InputManager lib.InputManager
	Entities     []lib.Entity
}

func NewFroggoGame() *FroggoGame {

	g := &FroggoGame{}
	g.Entities = []lib.Entity{}

	// add managers
	g.InputManager = &lib.EbitenInputManager{}

	// add objects
	player := game.NewPlayer(g.InputManager)
	g.Entities = append(g.Entities, player)

	return g
}

func (g *FroggoGame) Update() error {

	for _, v := range g.Entities {
		if err := v.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *FroggoGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for _, v := range g.Entities {
		v.Draw(screen)
	}
}

func (g *FroggoGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	froggoGame := NewFroggoGame()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("FROGGO - " + VERSION)
	if err := ebiten.RunGame(froggoGame); err != nil {
		log.Fatal(err)
	}
}
