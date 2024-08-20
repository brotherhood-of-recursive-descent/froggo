package game

import (
	"froggo/lib"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const RectangleWidthAndHeight = 256

type PlayerRectangle struct {
	Player *Player
	img    *ebiten.Image
	color  color.Color
}

var colorGreen = color.RGBA{25, 175, 25, 255}

// NewPlayerRectangle constructs a new player rectangle
func NewPlayerRectangle(p *Player) *PlayerRectangle {

	r := PlayerRectangle{
		Player: p,
		color:  colorGreen,
	}

	r.img = ebiten.NewImage(RectangleWidthAndHeight, RectangleWidthAndHeight)
	r.img.Fill(colorGreen)

	return &r
}

func (p *PlayerRectangle) IsActive() bool {
	return true // FIXME: for now always active
}

func (p *PlayerRectangle) Update() error {
	return nil
}

func (p *PlayerRectangle) Draw(screen *ebiten.Image) {

	x, y := derivePosition(p.Player.Position)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(p.img, op)
}

func derivePosition(p lib.Vec2) (float64, float64) {
	return p.X, p.Y
}
