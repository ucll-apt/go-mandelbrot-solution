package main

import (
	"encoding/base64"
	"encoding/binary"
	"log"
	"os"
)

func ExportBinary(path string, mandelbrots []Mandelbrot) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Couldn't open file")
	}
	defer file.Close()

	for _, mandelbrot := range mandelbrots {
		width := mandelbrot.Width()
		height := mandelbrot.Height()
		buffer := make([]byte, 8+3*width*height)
		binary.LittleEndian.PutUint32(buffer[0:], uint32(width))
		binary.LittleEndian.PutUint32(buffer[4:], uint32(height))

		i := 8
		for y := 0; y != height; y++ {
			for x := 0; x != width; x++ {
				c := byte(mandelbrot.Buffer[y][x])
				buffer[i] = c
				i++
				buffer[i] = c
				i++
				buffer[i] = c
				i++
			}
		}

		file.Write(buffer)
	}

	binary.Write(file, binary.LittleEndian, int32(0))
}

func ExportText(path string, mandelbrots []Mandelbrot) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Couldn't open file")
	}
	defer file.Close()

	for _, mandelbrot := range mandelbrots {
		file.WriteString("<<<")
		width := mandelbrot.Width()
		height := mandelbrot.Height()
		buffer := make([]byte, 8+3*width*height)
		binary.LittleEndian.PutUint32(buffer[0:], uint32(width))
		binary.LittleEndian.PutUint32(buffer[4:], uint32(height))

		i := 8
		for y := 0; y != height; y++ {
			for x := 0; x != width; x++ {
				c := byte(mandelbrot.Buffer[y][x])
				buffer[i] = c
				i++
				buffer[i] = c
				i++
				buffer[i] = c
				i++
			}
		}

		b64 := base64.StdEncoding.EncodeToString(buffer)
		file.WriteString(b64)
		file.WriteString(">>>")
	}
}
