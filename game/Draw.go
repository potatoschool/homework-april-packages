package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (view *Game) Draw(screen *ebiten.Image) {
	view.shape.Render(screen)
}
