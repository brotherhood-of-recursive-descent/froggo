package main

import (
	"froggo/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// to be replaced by build script
var VERSION = "latest"

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("FROGGO - " + VERSION)
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}
