package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

type instruction struct {
	code     int
	arg1mode int
	arg2mode int
	arg3mode int
}

func parseOpcode(code int) instruction {
	opcode := instruction{}
	digit := 0

	digit = code % 10
	code = code / 10
	opcode.code = ((code % 10) * 10) + digit
	code = code / 10

	opcode.arg1mode = code % 10
	code = code / 10

	opcode.arg2mode = code % 10
	code = code / 10

	opcode.arg3mode = code % 10

	return opcode
}

func runProgram(program []int) {
	cursor := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		opcode := parseOpcode(program[cursor])
		switch opcode.code {
		case 1:
			var arg1, arg2, resultDest int

			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			resultDest = program[cursor+3]

			program[resultDest] = arg1 + arg2
			cursor += 4
		case 2:
			var arg1, arg2, resultDest int

			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			resultDest = program[cursor+3]

			program[resultDest] = arg1 * arg2
			cursor += 4
		case 3:
			value := 0
			arg1 := program[cursor+1]

			fmt.Print("Input: ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			value = toInt(text)

			program[arg1] = value
			cursor += 2
		case 4:
			arg1 := program[cursor+1]
			fmt.Printf("Output: %d\n", program[arg1])
			cursor += 2
		case 5:
			var arg1, arg2 int
			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			if arg1 != 0 {
				cursor = arg2
			} else {
				cursor += 3
			}
		case 6:
			var arg1, arg2 int
			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			if arg1 == 0 {
				cursor = arg2
			} else {
				cursor += 3
			}
		case 7:
			var arg1, arg2, resultDest int
			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			resultDest = program[cursor+3]

			if arg1 < arg2 {
				program[resultDest] = 1
			} else {
				program[resultDest] = 0
			}

			cursor += 4
		case 8:
			var arg1, arg2, resultDest int
			if opcode.arg1mode == 1 {
				arg1 = program[cursor+1]
			} else {
				arg1 = program[program[cursor+1]]
			}

			if opcode.arg2mode == 1 {
				arg2 = program[cursor+2]
			} else {
				arg2 = program[program[cursor+2]]
			}

			resultDest = program[cursor+3]

			if arg1 == arg2 {
				program[resultDest] = 1
			} else {
				program[resultDest] = 0
			}
			cursor += 4
		case 99:
			return
		default:
			panic("Invalid OpCode")
		}
	}
}

func main() {
	fileString := readFile("input")
	strArray := strings.Split(string(fileString), ",")
	program := make([]int, len(strArray))

	for i, s := range strArray {
		program[i] = toInt(s)
	}

	fmt.Println("=== Intcode computer 2 ===")
	runProgram(program)
}
