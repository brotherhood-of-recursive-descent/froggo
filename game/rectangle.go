package game

import (
	"froggo/lib"

	"github.com/hajimehoshi/ebiten/v2"
)

const RectangleWidthAndHeight = 256
const BorderWidth = 1

type PlayerRectangle struct {
	Player *Player
	img    *ebiten.Image
}

// NewPlayerRectangle constructs a new player rectangle
func NewPlayerRectangle(p *Player) *PlayerRectangle {

	r := PlayerRectangle{
		Player: p,
	}

	grayBox := ebiten.NewImage(RectangleWidthAndHeight, RectangleWidthAndHeight)
	greenBox := ebiten.NewImage(RectangleWidthAndHeight-BorderWidth*2, RectangleWidthAndHeight-BorderWidth*2)
	grayBox.Fill(lib.ColorGray)
	greenBox.Fill(lib.ColorGreen)

	r.img = ebiten.NewImage(RectangleWidthAndHeight, RectangleWidthAndHeight)
	r.img.DrawImage(grayBox, nil)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(BorderWidth, BorderWidth)
	r.img.DrawImage(greenBox, &op)

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
