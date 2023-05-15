package shapes

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type IShape interface {
	Area() float64
	Perimeter() float64
	GetName() string
	Render(*ebiten.Image)
}

func PrintInfo(s IShape) {
	fmt.Println("--- Shape info ---")
	fmt.Printf("Shape name: %s\n", s.GetName())
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
