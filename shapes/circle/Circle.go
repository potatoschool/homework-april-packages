package circle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/shapes/shapes"
)

type Circle struct {
	Radius float64
	name   string
}

func New() Circle {
	var rInput string
	var r float64
	var err error

	for r == 0 {
		fmt.Println("Enter radius (number):")
		fmt.Scanln(&rInput)

		r, err = strconv.ParseFloat(rInput, 64)
		if err != nil {
			fmt.Println("Radius must be a number.")
		}
	}

	circle := Circle{r, "circle"}

	return circle
}

func ToShape() shapes.IShape {
	return New()
}
