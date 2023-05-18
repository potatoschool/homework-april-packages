package triangle

import (
	"fmt"

	"github.com/potatoschool/shapes/logger"
)

func (r Triangle) Perimeter() float64 {
	return r.Base + r.Left + r.Right
}

func (t *Triangle) calculatePerimeter() {
	t.perimeter = t.Perimeter()

	logger.Log(fmt.Sprintf("Calculated perimeter for triangle: %f", t.perimeter))
}
