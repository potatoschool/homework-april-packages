package triangle

import (
	"fmt"
	"strconv"

	"github.com/potatoschool/shapes/logger"
	"github.com/potatoschool/shapes/shapes"
)

type Triangle struct {
	ACorner        float64
	BCorner        float64
	BaseSideWidth  float64
	RightSideWidth float64
	LeftSideWidth  float64
	height         float64
	area           float64
	perimeter      float64
	name           string
}

func New() Triangle {
	var acInput, bcInput string
	var ACorner, BCorner, CCorner float64
	var err error
	var tri Triangle

	logger.Log("Triangle is chosen.")

	for CCorner == 0 {
		for ACorner == 0 {
			fmt.Println("Enter first corner angle in degrees (number):")
			fmt.Scanln(&acInput)

			ACorner, err = strconv.ParseFloat(acInput, 64)
			if err != nil {
				fmt.Println("Corner angle must be a number.")
			}
		}

		for BCorner == 0 {
			fmt.Println("Enter second corner angle in degrees (number):")
			fmt.Scanln(&bcInput)

			BCorner, err = strconv.ParseFloat(bcInput, 64)
			if err != nil {
				fmt.Println("Corner angle must be a number.")
			}
		}

		CCorner = 180 - ACorner - BCorner

		if CCorner <= 0 {
			fmt.Println("Corners sum must be equal or less then 180.")
			ACorner = 0
			BCorner = 0
			CCorner = 0
		} else {
			tri = Triangle{ACorner, BCorner, 100, 100, 100, 0, 0, 0, "triangle"}
		}
	}

	logger.Log(
		fmt.Sprintf(
			"Triangle initialized with A corner angle %f, B corner angle %f and Calculated Top corner angle %f",
			ACorner,
			BCorner,
			CCorner,
		),
	)

	tri.calculateArea()
	tri.calculatePerimeter()
	tri.calculateHeight()

	return tri
}

func ToShape() shapes.IShape {
	return New()
}
