package player

import (
	"froggo/game"
	"froggo/lib"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerRectangle struct {
	game.Entity

	Position lib.Vec2
	img      *ebiten.Image
}

const RectangleWidth = 256 // for now here; square cells
const RectangleHeight = 256

var ColorGreen = color.RGBA{25, 175, 25, 255} // TODO move to lib

func NewPlayerRectangle() *PlayerRectangle {

	r := PlayerRectangle{}

	r.Position.X = 100 // FIXME: middle, bottom of the screen
	r.Position.Y = 100

	r.img = ebiten.NewImage(RectangleHeight, RectangleWidth)

	return &r
}

func (p *PlayerRectangle) IsActive() bool {
	return true // FIXME: for now always active
}

func (p *PlayerRectangle) Update() error {
	return nil // TODO: update position if got movement component
}

func (p *PlayerRectangle) Draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	p.img.Fill(ColorGreen)
	screen.DrawImage(p.img, nil) // TODO: use proper options
}
