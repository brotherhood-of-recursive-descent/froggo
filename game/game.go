package game

import (
	"froggo/game/player"
	"froggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game contains all global available objects
type Game struct {
	InputManager lib.InputManager
	Entities     []lib.Entity
}

func NewGame() *Game {

	g := &Game{}
	g.Entities = []lib.Entity{}

	// add managers
	g.InputManager = &lib.EbitenInputManager{}

	// add objects
	player := player.NewPlayer(g.InputManager)
	g.Entities = append(g.Entities, player)

	return g
}

func (g *Game) Update() error {

	for _, v := range g.Entities {
		if err := v.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	for _, v := range g.Entities {
		v.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
