package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (view *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x23, 0x23, 0x23, 0xff})
	view.shape.Render(screen)
}
