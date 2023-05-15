package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/potatoschool/shapes/game"
	"github.com/potatoschool/shapes/shapes"
	"github.com/potatoschool/shapes/shapes/circle"
	"github.com/potatoschool/shapes/shapes/rectangle"
)

func main() {
	var i int
	var shapeName string
	var shape shapes.IShape
	var availableShapesAllAliases []string
	var availableShapesMainAliases []string
	var launchVisualUserInput string
	var launchVisual bool

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

	fmt.Println("--- \033[33mShapes Builder\033[0m --- ")
	fmt.Printf("\033[32mAvailable shapes:\033[0m \n%s\n", strings.Join(availableShapesMainAliases, "\n"))
	fmt.Println("\033[32mAvailable commands:\033[0m")
	fmt.Println("\033[33m--list:\033[0m to list all available shapes and it's aliases")
	fmt.Println("--------------------------")

	for shape == nil {
		fmt.Println("Enter command")
		fmt.Scanln(&shapeName)

		shapeName = strings.ToLower(shapeName)
		targetShape, exists := availableShapes[shapeName]

		if exists {
			shape = targetShape()
		} else if shapeName == "--list" {
			fmt.Println(strings.Join(availableShapesAllAliases, ""))
		} else {
			fmt.Printf("There is no available shape with name %s\n", shapeName)
		}
	}

	shapes.PrintInfo(shape)

	for launchVisualUserInput == "" {
		fmt.Println("Render the shape in 2D? (Y/n)")
		fmt.Scanln(&launchVisualUserInput)

		if launchVisualUserInput == "Y" || launchVisualUserInput == "y" {
			launchVisual = true
		} else if launchVisualUserInput == "N" || launchVisualUserInput == "n" {
			launchVisual = false
		} else {
			launchVisualUserInput = ""
		}
	}

	if !launchVisual {
		os.Exit(0)
		return
	}

	settings := game.Settings{}

	settings.SetHeight(320)
	settings.SetWidth(320)

	fmt.Println("Launching render...")

	if err := game.Render(shape, settings); err != nil {
		fmt.Println("Can't launch render.")
		fmt.Printf("---\nError:\n\n%s\n---", err.Error())
		os.Exit(1)
		return
	}
}
