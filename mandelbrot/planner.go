package main

type Job func()

type Planner interface {
	jobCount() int
	getJob(index int) Job
}

type PixelPlanner struct {
	mandelbrots []Mandelbrot
}

func (p PixelPlanner) jobCount() int {
	nfractals := len(p.mandelbrots)
	width := p.mandelbrots[0].Width()
	height := p.mandelbrots[0].Height()
	return nfractals * width * height
}

func (p PixelPlanner) getJob(index int) Job {
	width := p.mandelbrots[0].Width()
	height := p.mandelbrots[0].Height()
	pixelCount := width * height

	frameIndex := index / pixelCount
	pixelIndex := index % pixelCount
	x := pixelIndex % width
	y := pixelIndex / width

	return func() {
		p.mandelbrots[frameIndex].ComputeSingle(x, y)
	}
}

type RowPlanner struct {
	mandelbrots []Mandelbrot
}

func (p RowPlanner) jobCount() int {
	nfractals := len(p.mandelbrots)
	height := p.mandelbrots[0].Height()
	return nfractals * height
}

func (p RowPlanner) getJob(index int) Job {
	height := p.mandelbrots[0].Height()
	rowIndex := index % height
	frameIndex := index / height

	return func() {
		p.mandelbrots[frameIndex].ComputeRow(rowIndex)
	}
}

type FramePlanner struct {
	mandelbrots []Mandelbrot
}

func (p FramePlanner) jobCount() int {
	return len(p.mandelbrots)
}

func (p FramePlanner) getJob(index int) Job {
	return func() {
		p.mandelbrots[index].ComputeAll()
	}
}
