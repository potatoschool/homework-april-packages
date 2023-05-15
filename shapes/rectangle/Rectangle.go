package rectangle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/shapes/logger"
	"github.com/potatoschool/shapes/shapes"
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

	logger.Log("Rectangle is chosen.")

	for w == 0 {
		fmt.Println("Enter width (number):")
		fmt.Scanln(&wInput)

		w, err = strconv.ParseFloat(wInput, 64)
		if err != nil {
			fmt.Println("Width must be a number.")
		}
	}

	for h == 0 {
		fmt.Println("Enter height (number):")
		fmt.Scanln(&hInput)

		h, err = strconv.ParseFloat(hInput, 64)
		if err != nil {
			fmt.Println("Height must be a number.")
		}
	}

	rect := Rectangle{w, h, "rectangle"}

	logger.Log(fmt.Sprintf("Rectangle initialized with width %f and height %f", w, h))

	return rect
}

func ToShape() shapes.IShape {
	return New()
}
