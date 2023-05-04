package rectangle

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
