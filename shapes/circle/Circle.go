package circle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/shapes/logger"
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

	logger.Log("Circle is chosen.")

	for r == 0 {
		fmt.Println("Enter radius (number):")
		fmt.Scanln(&rInput)

		r, err = strconv.ParseFloat(rInput, 64)
		if err != nil {
			fmt.Println("Radius must be a number.")
		}
	}

	circle := Circle{r, "circle"}

	logger.Log(fmt.Sprintf("Circle initialized with radius %f", r))

	return circle
}

func ToShape() shapes.IShape {
	return New()
}
