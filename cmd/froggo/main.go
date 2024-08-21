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
	world    game.World
	Entities []lib.Entity
}

func NewFroggoGame() *FroggoGame {

	g := &FroggoGame{}
	g.Entities = []lib.Entity{}

	g.world = game.World{}
	g.world.InputManager = &lib.EbitenInputManager{}
	screenWidth, screenHeight := ebiten.Monitor().Size() // does not work on mobiles
	g.world.ScreenDimension.X, g.world.ScreenDimension.Y = float64(screenWidth), float64(screenHeight)

	// add objects
	player := game.NewPlayer(g.world)
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

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("FROGGO - " + VERSION)

	froggoGame := NewFroggoGame()
	if err := ebiten.RunGame(froggoGame); err != nil {
		log.Fatal(err)
	}
}
