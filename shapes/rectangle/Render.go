package rectangle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (r Rectangle) Render(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(screen.Bounds().Dx()/2)-float32(r.Width/2),
		float32(screen.Bounds().Dy()/2)-float32(r.Height/2),
		float32(r.Width),
		float32(r.Height),
		color.NRGBA{0xbb, 0x00, 0x00, 0xf0},
		false,
	)
}
