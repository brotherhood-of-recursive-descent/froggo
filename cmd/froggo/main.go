package main

import (
	"froggo/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// to be replaced by build script
var VERSION = "latest"

func main() {

	froggoGame := game.NewGame()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("FROGGO - " + VERSION)
	if err := ebiten.RunGame(froggoGame); err != nil {
		log.Fatal(err)
	}
}
