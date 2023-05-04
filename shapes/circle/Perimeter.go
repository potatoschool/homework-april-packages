package circle

import "math"

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
