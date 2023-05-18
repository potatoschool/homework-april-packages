package triangle

import (
	"fmt"
	"math"

	"github.com/potatoschool/shapes/logger"
)

func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2

	area := math.Sqrt(s * (s - t.BaseSideWidth) * (s - t.LeftSideWidth) * (s - t.RightSideWidth))

	return area
}

func (t *Triangle) calculateArea() {
	t.area = t.Area()

	logger.Log(fmt.Sprintf("Calculated area for triangle: %f", t.area))
}
