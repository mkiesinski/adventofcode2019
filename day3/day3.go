package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type point struct {
	x     int
	y     int
	steps int
}
type pointSlice []point

type intsec struct {
	point1 point
	point2 point
}
type intsecSlice []intsec

func (p *point) manhattanDistance() int {
	return intAbs(p.x) + intAbs(p.y)
}

func (p *point) applyVector(v point) {
	p.x += v.x
	p.y += v.y
}

func (p *point) equals(c point) bool {
	return (p.x == c.x) && (p.y == c.y)
}

type wire struct {
	points pointSlice
}

func (w *wire) findPoint(p point) (point, bool) {
	for i := range w.points {
		if w.points[i].equals(p) {
			return w.points[i], true
		}
	}
	return point{0, 0, 0}, false
}

func generateWire(instString string) wire {
	points := pointSlice{}
	current := point{0, 0, 0}

	instructions := strings.Split(string(instString), ",")

	for i := range instructions {
		direction := string(instructions[i][:1])
		distance := toInt(string(instructions[i][1:]))

		vector := point{0, 0, 0}
		switch direction {
		case "U":
			vector.y = 1
		case "D":
			vector.y = -1
		case "L":
			vector.x = -1
		case "R":
			vector.x = 1
		}

		for j := 0; j < distance; j++ {
			current.applyVector(vector)
			current.steps++
			points = append(points, current)
		}
	}

	generated := wire{points}

	return generated
}

func getIntersections(w1 wire, w2 wire) intsecSlice {
	intersections := intsecSlice{}

	for i := range w1.points {
		ipoint, found := w2.findPoint(w1.points[i])
		if found {
			intersections = append(intersections, intsec{w1.points[i], ipoint})
		}
	}
	return intersections
}

func findClosest(list intsecSlice) int {
	min := list[0].point1.manhattanDistance()

	for i := range list {
		if list[i].point1.manhattanDistance() < min {
			min = list[i].point1.manhattanDistance()
		}
	}

	return min
}

func findFastest(list intsecSlice) int {
	min := list[0].point1.steps + list[0].point2.steps

	for i := range list {
		distance := list[i].point1.steps + list[i].point2.steps
		if distance < min {
			min = distance
		}
	}

	return min
}

func main() {
	fileString := readFile("input")
	inputArray := strings.Split(strings.Replace(string(fileString), "\r\n", "\n", -1), "\n")

	w1 := generateWire(inputArray[0])
	w2 := generateWire(inputArray[1])

	intersections := getIntersections(w1, w2)

	fmt.Println("=== Part 1 ===")
	fmt.Printf("Intersection distance closest to the point of origin: %d\n", findClosest(intersections))
	fmt.Println("=== Part 2 ===")
	fmt.Printf("Intersection reached fastest by the signal: %d\n", findFastest(intersections))
}
