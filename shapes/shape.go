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

func GetInfo(s IShape) {
	fmt.Println("--- Результат ---")
	fmt.Printf("Фигура: %s\n", s.GetName())
	fmt.Printf("Площадь: %.2f\n", s.Area())
	fmt.Printf("Периметр: %.2f\n", s.Perimeter())
}
