package triangle

import (
	"fmt"

	"github.com/potatoschool/shapes/logger"
)

func (r Triangle) Perimeter() float64 {
	return r.BaseSideWidth + r.LeftSideWidth + r.RightSideWidth
}

func (t *Triangle) calculatePerimeter() {
	t.perimeter = t.Perimeter()

	logger.Log(fmt.Sprintf("Calculated perimeter for triangle: %f", t.perimeter))
}
