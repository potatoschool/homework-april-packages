package triangle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/shapes/logger"
	"github.com/potatoschool/shapes/shapes"
)

type Triangle struct {
	Base      float64
	Right     float64
	Left      float64
	height    float64
	area      float64
	perimeter float64
	name      string
}

func New() Triangle {
	var bInput, rInput, lInput string
	var b, r, l float64
	var err error

	logger.Log("Triangle is chosen.")

	for b == 0 {
		fmt.Println("Enter base (number):")
		fmt.Scanln(&bInput)

		b, err = strconv.ParseFloat(bInput, 64)
		if err != nil {
			fmt.Println("Width must be a number.")
		}
	}

	for r == 0 {
		fmt.Println("Enter right width (number):")
		fmt.Scanln(&rInput)

		r, err = strconv.ParseFloat(rInput, 64)
		if err != nil {
			fmt.Println("Right Width must be a number.")
		}
	}

	for l == 0 {
		fmt.Println("Enter left width (number):")
		fmt.Scanln(&lInput)

		l, err = strconv.ParseFloat(lInput, 64)
		if err != nil {
			fmt.Println("Left Width must be a number.")
		}
	}

	tri := Triangle{b, r, l, 0, 0, 0, "triangle"}

	logger.Log(fmt.Sprintf("Triangle initialized with base %f, right %f and left %f", b, r, l))

	tri.calculateArea()
	tri.calculatePerimeter()
	tri.calculateHeight()

	return tri
}

func ToShape() shapes.IShape {
	return New()
}
