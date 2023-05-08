package packages

import (
	"fmt"
	"strings"

	"github.com/potatoschool/homework-april-packages/shapes"
	"github.com/potatoschool/homework-april-packages/shapes/circle"
	"github.com/potatoschool/homework-april-packages/shapes/rectangle"
)

func Load() {
	var shapeName string
	var shape shapes.IShape

	for shape == nil {
		fmt.Println("Выберите фигуру (rectangle или circle)")
		fmt.Scanln(&shapeName)

		shapeName = strings.ToLower(shapeName)

		switch shapeName {
		case "Rectangle":
			fallthrough
		case "r":
			fallthrough
		case "R":
			fallthrough
		case "rectangle":
			shape = rectangle.New()
		case "Circle":
			fallthrough
		case "c":
			fallthrough
		case "circle":
			shape = circle.New()
		default:
			fmt.Println("Доступны только rectangle или circle")
		}
	}

	shapes.GetInfo(shape)
}
