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

	fmt.Println("=== Part 1 ===")
	inout := make(chan int)
	value := 0
	prog1 := intcode.Program{Memory: append(program[:0:0], program...), ChIn: inout, ChOut: inout}
	go func() {
		prog1.Run()
		close(inout)
	}()
	inout <- 1

	for read := range inout {
		value = read
	}

	fmt.Printf("Ouput for input 1 : %d\n", value)

	fmt.Println("=== Part 2 ===")
	value = 0
	inout = make(chan int)
	prog2 := intcode.Program{Memory: append(program[:0:0], program...), ChIn: inout, ChOut: inout}
	go func() {
		prog2.Run()
		close(inout)
	}()
	inout <- 5

	for read := range inout {
		value = read
	}

	fmt.Printf("Ouput for input 5 : %d\n", value)
}
