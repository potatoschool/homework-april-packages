package circle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/homework-april-packages/shapes"
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
		fmt.Println("Введите радиус (число):")
		fmt.Scanln(&rInput)

		r, err = strconv.ParseFloat(rInput, 64)
		if err != nil {
			fmt.Println("Радиус должен быть числом.")
		}
	}

	circle := Circle{r, "circle"}

	return circle
}

func ToShape() shapes.IShape {
	return New()
}
