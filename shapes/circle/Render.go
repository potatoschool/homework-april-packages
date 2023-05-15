package circle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (c Circle) Render(screen *ebiten.Image) {
	vector.DrawFilledCircle(
		screen,
		float32(screen.Bounds().Dx()/2),
		float32(screen.Bounds().Dy()/2),
		float32(c.Radius),
		color.NRGBA{0xbb, 0x00, 0x00, 0xf0},
		false,
	)
}
