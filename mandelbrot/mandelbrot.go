package main

import "math/cmplx"

type Mandelbrot struct {
	Rectangle         Rectangle
	MaximumIterations int
	MaximumMagnitude  float64
	Buffer            [][]int
}

func NewMandelbrot(hres, vres int, rect *Rectangle, maxIterations int, maxMagnitude float64) *Mandelbrot {
	buffer := make([][]int, vres)

	for i := 0; i != vres; i++ {
		buffer[i] = make([]int, hres)
	}

	return &Mandelbrot{
		Rectangle:         *rect,
		MaximumIterations: maxIterations,
		MaximumMagnitude:  maxMagnitude,
		Buffer:            buffer,
	}
}

func (m *Mandelbrot) Width() int {
	return len(m.Buffer[0])
}

func (m *Mandelbrot) Height() int {
	return len(m.Buffer)
}

func (m *Mandelbrot) computeInitialValue(x, y int) complex128 {
	p := Point{X: float64(x) / float64(m.Width()), Y: float64(y) / float64(m.Height())}
	q := m.Rectangle.FromRelative(&p)

	return complex(q.X, q.Y)
}

func (m *Mandelbrot) ComputeSingle(x, y int) {
	c := m.computeInitialValue(x, y)
	z := c
	i := 0

	for cmplx.Abs(z) < m.MaximumMagnitude && i < m.MaximumIterations {
		z = z*z + c
		i++
	}

	m.Buffer[y][x] = i
}

func (m *Mandelbrot) ComputeRow(y int) {
	for x := 0; x != m.Width(); x++ {
		m.ComputeSingle(x, y)
	}
}

func (m *Mandelbrot) ComputeAll() {
	for y := 0; y != m.Height(); y++ {
		m.ComputeRow(y)
	}
}
