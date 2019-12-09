package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/mkiesinski/adventofcode2019/intcode"
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

func main() {
	fileString := readFile("input")
	strArray := strings.Split(string(fileString), ",")
	program := make([]int, len(strArray))

	for i, s := range strArray {
		program[i] = toInt(s)
	}

	result := 0

	fmt.Println("=== Part 1 ===")
	inout := make(chan int)
	prog := intcode.Program{ChIn: inout, ChOut: inout}
	prog.LoadMemory(program)

	go func() {
		prog.Run()
		close(inout)
	}()
	inout <- 1

	for out := range inout {
		result = out
	}

	fmt.Printf("Control code of BOOST program: %d\n", result)

	fmt.Println("=== Part 2 ===")
	inout = make(chan int)
	prog2 := intcode.Program{ChIn: inout, ChOut: inout}
	prog2.LoadMemory(program)

	go func() {
		prog2.Run()
		close(inout)
	}()
	inout <- 2

	for out := range inout {
		result = out
	}

	fmt.Printf("Distress signal coordinates: %d\n", result)
}
