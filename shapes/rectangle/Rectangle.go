package rectangle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/homework-april-packages/shapes"
)

type Rectangle struct {
	Width  float64
	Height float64
	name   string
}

func New() Rectangle {
	var wInput, hInput string
	var w, h float64
	var err error

	for w == 0 {
		fmt.Println("Введите ширину (число):")
		fmt.Scanln(&wInput)

		w, err = strconv.ParseFloat(wInput, 64)
		if err != nil {
			fmt.Println("Ширина должна быть числом.")
		}
	}

	for h == 0 {
		fmt.Println("Введите длину (число):")
		fmt.Scanln(&hInput)

		h, err = strconv.ParseFloat(hInput, 64)
		if err != nil {
			fmt.Println("Ширина должна быть числом.")
		}
	}

	rect := Rectangle{w, h, "rectangle"}

	return rect
}

func ToShape() shapes.IShape {
	return New()
}
