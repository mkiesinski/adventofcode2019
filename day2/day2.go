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

func findNounVerb(baseProgram []int, search int) int {
	var noun, verb int

	for verb = 0; verb < 100; verb++ {
		for noun = 0; noun < 100; noun++ {
			program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: nil, ChOut: nil}
			program.Memory[1] = noun
			program.Memory[2] = verb
			program.Run()
			value := program.Memory[0]
			if value == search {
				return (100 * noun) + verb
			}
		}
	}

	return 0
}

func main() {
	fileString := readFile("input")
	strArray := strings.Split(string(fileString), ",")
	program := make([]int, len(strArray))

	for i, s := range strArray {
		program[i] = toInt(s)
	}

	fmt.Println("=== Part 1 ===")
	prog := intcode.Program{Memory: append(program[:0:0], program...), ChIn: nil, ChOut: nil}
	prog.Memory[1] = 12
	prog.Memory[2] = 2
	prog.Run()
	fmt.Printf("1202 error execution result: %d\n", prog.Memory[0])

	fmt.Println("=== Part 2 ===")
	result := 19690720
	fmt.Printf("Noun and verb producing rersult %d: %d", result, findNounVerb(append(program[:0:0], program...), result))
}
