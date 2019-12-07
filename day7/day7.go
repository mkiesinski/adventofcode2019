package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type instruction struct {
	code     int
	arg1mode int
	arg2mode int
	arg3mode int
}

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

func runProgram(program []int, in chan int, out chan int) {
	cursor := 0
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
			value := <-in
			arg1 := program[cursor+1]
			program[arg1] = value
			cursor += 2
		case 4:
			arg1 := program[cursor+1]
			out <- program[arg1]
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

func verifySequence(sequence string) bool {
	for i, s := range sequence {
		if toInt(string(s)) > 4 {
			return false
		}
		for j, s2 := range sequence {
			if s == s2 && i != j {
				return false
			}
		}
	}
	return true
}

func verifyFeedbackSequence(sequence string) bool {
	for i, s := range sequence {
		if toInt(string(s)) < 5 {
			return false
		}
		for j, s2 := range sequence {
			if s == s2 && i != j {
				return false
			}
		}
	}
	return true
}

func calculateSequence(baseProgram []int, sequence string) int {
	inout := make(chan int)
	value := 0

	if !verifySequence(sequence) {
		return value
	}

	for _, s := range sequence {
		setting := toInt(string(s))
		go runProgram(append(baseProgram[:0:0], baseProgram...), inout, inout)
		inout <- setting
		inout <- value
		value = <-inout
	}

	return value
}

func calculateFeedbackLoop(baseProgram []int, sequence string) int {
	value := 0
	if !verifyFeedbackSequence(sequence) {
		return value
	}

	var waitgroup sync.WaitGroup
	var controlgroup sync.WaitGroup

	inA := make(chan int)
	inB := make(chan int)
	inC := make(chan int)
	inD := make(chan int)
	inE := make(chan int)
	feedBack := make(chan int)

	waitgroup.Add(5)
	controlgroup.Add(1)
	go func() {
		runProgram(append(baseProgram[:0:0], baseProgram...), inA, inB)
		waitgroup.Done()
	}()
	go func() {
		runProgram(append(baseProgram[:0:0], baseProgram...), inB, inC)
		waitgroup.Done()
	}()
	go func() {
		runProgram(append(baseProgram[:0:0], baseProgram...), inC, inD)
		waitgroup.Done()
	}()
	go func() {
		runProgram(append(baseProgram[:0:0], baseProgram...), inD, inE)
		waitgroup.Done()
	}()
	go func() {
		runProgram(append(baseProgram[:0:0], baseProgram...), inE, feedBack)
		close(feedBack)
		waitgroup.Done()
	}()
	go func() {
		for elem := range feedBack {
			inA <- elem
		}
		controlgroup.Done()
	}()

	inA <- toInt(string(sequence[0]))
	inB <- toInt(string(sequence[1]))
	inC <- toInt(string(sequence[2]))
	inD <- toInt(string(sequence[3]))
	inE <- toInt(string(sequence[4]))
	inA <- 0

	waitgroup.Wait()
	result := <-inA

	controlgroup.Wait()
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
	max := 0
	for i := 0; i <= 99999; i++ {
		calculated := calculateSequence(program, fmt.Sprintf("%05d", i))
		if calculated > max {
			max = calculated
		}
	}
	fmt.Printf("Highest signal sent to thrusters: %d\n", max)

	fmt.Println("=== Part 2 ===")
	max = 0
	for i := 0; i <= 99999; i++ {
		calculated := calculateFeedbackLoop(program, fmt.Sprintf("%05d", i))
		if calculated > max {
			max = calculated
		}
	}
	fmt.Printf("Highest signal sent to thrusters with feedback loop: %d\n", max)
}
