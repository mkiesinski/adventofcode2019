package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	imageWidth  = 25
	imageHeight = 6
	pixels      = imageWidth * imageHeight
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func calculateLayer(layer string) (int, int) {
	count := make([]int, 3)

	for _, s := range layer {
		digit := s - 48
		count[digit]++
	}

	return count[0], count[1] * count[2]
}

func applyLayer(canvas []int, layer string) {
	for i, s := range layer {
		if canvas[i] == 2 {
			canvas[i] = toInt(string(s))
		}
	}
}

func printCanvas(canvas []int) {
	for i, pixel := range canvas {
		if i%(imageWidth) == 0 {
			fmt.Printf("\n")
		}
		if pixel == 1 {
			fmt.Print("▮")
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	fileString := readFile("input")
	layers := make([]string, 0)
	layerCount := len(fileString) / pixels

	for i := 0; i < layerCount; i++ {
		layers = append(layers, fileString[i*pixels:i*pixels+pixels])
	}

	fmt.Println("=== Part 1 ===")
	var maxZeroCount, result int
	maxZeroCount = pixels

	for _, layer := range layers {
		zeroCount, res := calculateLayer(layer)

		if zeroCount < maxZeroCount {
			maxZeroCount = zeroCount
			result = res
		}
	}

	fmt.Printf("Layers control sum: %d\n", result)

	fmt.Println("=== Part 2 ===")
	canvas := make([]int, pixels)
	for i := range canvas {
		canvas[i] = 2
	}

	for _, layer := range layers {
		applyLayer(canvas, layer)
	}

	printCanvas(canvas)
	fmt.Println()
}
