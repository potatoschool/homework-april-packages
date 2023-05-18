package triangle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (t Triangle) Render(screen *ebiten.Image) {
	var path vector.Path

	path.MoveTo(0, float32(t.height))
	path.LineTo(float32(t.Base), float32(t.height))
	path.LineTo(float32(t.Base), 0)
	path.LineTo(0, float32(t.height))

	var canvas = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	canvas.Fill(color.RGBA{200, 0, 0, 255})

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	screen.DrawTriangles(vs, is, canvas, &ebiten.DrawTrianglesOptions{
		AntiAlias: false,
	})
}
