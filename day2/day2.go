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

func runProgram(program []int) int {
	for i := 0; i < len(program); i += 4 {
		switch program[i] {
		case 1:
			arg1 := program[i+1]
			arg2 := program[i+2]
			resultDest := program[i+3]
			program[resultDest] = program[arg1] + program[arg2]
		case 2:
			arg1 := program[i+1]
			arg2 := program[i+2]
			resultDest := program[i+3]
			program[resultDest] = program[arg1] * program[arg2]
		case 99:
			return program[0]
		default:
			panic("Invalid OpCode")
		}
	}

	return 0
}

func findNounVerb(baseProgram []int, search int) int {
	var noun, verb int

	for verb = 0; verb < 100; verb++ {
		for noun = 0; noun < 100; noun++ {
			program := append(baseProgram[:0:0], baseProgram...)
			program[1] = noun
			program[2] = verb
			value := runProgram(program)
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

	program[1] = 12
	program[2] = 2
	fmt.Println("=== Part 1 ===")
	fmt.Printf("1202 error execution result: %d\n", runProgram(append(program[:0:0], program...)))

	fmt.Println("=== Part 2 ===")
	result := 19690720
	fmt.Printf("Noun and verb producing rersult %d: %d", result, findNounVerb(append(program[:0:0], program...), result))
}
