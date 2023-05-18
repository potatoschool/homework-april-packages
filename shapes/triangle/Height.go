package triangle

import (
	"fmt"

	"github.com/potatoschool/shapes/logger"
)

func (t Triangle) Height() float64 {
	height := (2 * t.area) / t.BaseSideWidth

	fmt.Println(height)

	return height
}

func (t *Triangle) calculateHeight() {
	t.height = t.Height()

	logger.Log(fmt.Sprintf("Calculated height for triangle: %f", t.height))
}
