package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/potatoschool/homework-april-packages/shapes"
)

func Render(shape shapes.IShape, settings Settings) error {
	View := Game{
		shape:    shape,
		settings: settings,
	}

	return ebiten.RunGame(&View)
}
