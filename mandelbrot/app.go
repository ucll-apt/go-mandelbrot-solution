package main

import (
	"fmt"
	"time"
)

var initialWidth float64 = 3
var minimumWidth float64 = 0.0001
var zoomFactor float64 = 0.95
var x float64 = -0.761574
var y float64 = -0.0847596
var horizontalResolution int = 1920
var verticalResolution int = 1080
var maxIterations int = 255
var maxMagnitude float64 = 5.0

func createFrames() []Mandelbrot {
	result := []Mandelbrot{}
	ratio := float64(horizontalResolution) / float64(verticalResolution)
	width := initialWidth

	for width > minimumWidth {
		rectangle := NewRectangle(&Point{X: x, Y: y}, ratio, float64(width))
		mandelbrot := NewMandelbrot(horizontalResolution, verticalResolution, rectangle, maxIterations, maxMagnitude)
		result = append(result, *mandelbrot)
		width *= zoomFactor
	}

	return result
}

func main() {
	mandelbrots := createFrames()
	fmt.Printf("Rendering %d frames\n", len(mandelbrots))

	planner := RowPlanner{mandelbrots: mandelbrots}
	scheduler := ParallelScheduler{}

	before := time.Now()
	scheduler.Schedule(planner)

	elapsed := time.Since(before)

	fmt.Printf("Used %s seconds\n", elapsed)

	ExportText("g:/temp/gomandel.wif", mandelbrots)
}
