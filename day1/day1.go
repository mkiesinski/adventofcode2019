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

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateMoreFuel(mass int) int {
	var fuelAdded = calculateFuel(mass)
	var sum = fuelAdded

	for fuelAdded > 0 {
		fuelAdded = calculateFuel(fuelAdded)
		if fuelAdded < 1 {
			break
		}
		sum += fuelAdded
	}

	return sum
}

func main() {
	fileString := readFile("input")
	inputStr := strings.Split(strings.Replace(string(fileString), "\r\n", "\n", -1), "\n")
	input := make([]int, len(inputStr))

	for i, s := range inputStr {
		input[i] = toInt(string(s))
	}

	fmt.Println("=== Part 1 ===")
	sum := 0

	for _, mass := range input {
		sum += calculateFuel(mass)
	}
	fmt.Printf("Fuel required: %d\n", sum)

	fmt.Println("=== Part 2 ===")
	sum = 0

	for _, mass := range input {
		sum += calculateMoreFuel(mass)
	}
	fmt.Printf("Fuel required: %d\n", sum)
}
