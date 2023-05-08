package main

import (
	"fmt"
	"strings"

	"github.com/potatoschool/homework-april-packages/shapes"
	"github.com/potatoschool/homework-april-packages/shapes/circle"
	"github.com/potatoschool/homework-april-packages/shapes/rectangle"
)

func main() {
	var i int
	var shapeName string
	var shape shapes.IShape
	var availableShapesAllAliases []string
	var availableShapesMainAliases []string

	availableShapes := map[string]func() shapes.IShape{}

	ShapesData := map[string]func() shapes.IShape{
		rectangle.GetAlias(): rectangle.ToShape,
		circle.GetAlias():    circle.ToShape,
	}

	for shapeAliases, ShapeConstructor := range ShapesData {
		aliases := strings.Split(shapeAliases, ",")
		mainAlias := aliases[0]

		availableShapesMainAliases = append(availableShapesMainAliases, fmt.Sprintf("%d) %s", i+1, mainAlias))
		availableShapesAllAliases = append(availableShapesAllAliases, fmt.Sprintf("%d. %s:\n", i+1, mainAlias))

		for _, shapeAlias := range aliases {
			availableShapes[shapeAlias] = ShapeConstructor
			availableShapesAllAliases = append(availableShapesAllAliases, fmt.Sprintf(" - %s\n", shapeAlias))
		}

		availableShapes[fmt.Sprintf("%d", i+1)] = ShapeConstructor
		i++
	}

	fmt.Println("--- \033[33mКонструктор Фигур\033[0m --- ")
	fmt.Printf("\033[32mДоступные фигуры:\033[0m \n%s\n", strings.Join(availableShapesMainAliases, "\n"))
	fmt.Println("\033[32mДоступные команды:\033[0m")
	fmt.Println("\033[33m--list:\033[0m для просмотра всех доступных фигур и их алиасов")
	fmt.Println("--------------------------")

	for shape == nil {
		fmt.Println("Введите команду")
		fmt.Scanln(&shapeName)

		shapeName = strings.ToLower(shapeName)
		targetShape, exists := availableShapes[shapeName]

		if exists {
			shape = targetShape()
		} else if shapeName == "--list" {
			fmt.Println(strings.Join(availableShapesAllAliases, ""))
		} else {
			fmt.Println("Фигуры не существует!")
		}
	}

	shapes.GetInfo(shape)
}
