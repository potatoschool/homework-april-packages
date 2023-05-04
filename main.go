package packages

import (
	"fmt"
	"strings"

	"github.com/potatoschool/homework-april-packages/shapes"
	"github.com/potatoschool/homework-april-packages/shapes/rectangle"
)

func Load() {
	var shapeName string
	var shape shapes.IShape

	sh := map[string]func() shapes.IShape{
		"rectangle": rectangle.ToShape,
		"circle":    rectangle.ToShape,
	}

	for shape == nil {
		fmt.Println("Выберите фигуру (rectangle или circle)")
		fmt.Scanln(&shapeName)

		shapeName = strings.ToLower(shapeName)
		targetShape, exists := sh[shapeName]

		if exists {
			shape = targetShape()
		} else {
			fmt.Println("Доступны только rectangle или circle")
		}
	}

	shapes.GetInfo(shape)
}
