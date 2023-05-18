package triangle

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func getVertices(startX, startY, baseSideWidth, rightSideWidth, leftSideWidth float64, ACorner, BCorner float64) (float64, float64, float64, float64, float64, float64) {
	ACornerRad := ACorner * math.Pi / 180.0
	BCornerRad := BCorner * math.Pi / 180.0
	x0 := startX
	y0 := startY
	// the base is aligned with the x-axis
	x1 := x0 + baseSideWidth
	y1 := y0
	// height of the triangle
	h := rightSideWidth * math.Sin(BCornerRad)
	// the top vertex can be calculated using the height and the left bottom corner
	x2 := x0 + h/math.Tan(ACornerRad)
	y2 := y0 + h

	return x0, y0, x1, y1, x2, y2
}

func (t Triangle) Render(screen *ebiten.Image) {

	x0, y0, x1, y1, x2, y2 := getVertices(50, 50, t.BaseSideWidth, t.RightSideWidth, t.LeftSideWidth, t.ACorner, t.BCorner)

	var path vector.Path

	path.MoveTo(float32(x0), float32(y0)) // Start to Draw Base
	path.LineTo(float32(x1), float32(y1)) // Draw Base
	path.LineTo(float32(x2), float32(y2)) // Draw Right
	path.LineTo(float32(x0), float32(y0))

	var canvas = ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())
	canvas.Fill(color.RGBA{200, 0, 0, 255})

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	screen.DrawTriangles(vs, is, canvas, &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	})
}
