package circle

import "math"

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
