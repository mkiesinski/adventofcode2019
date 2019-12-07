package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"

	"github.com/mkiesinski/adventofcode2019/intcode"
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
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inout, ChOut: inout}
		go program.Run()
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
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inA, ChOut: inB}
		program.Run()
		waitgroup.Done()
	}()
	go func() {
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inB, ChOut: inC}
		program.Run()
		waitgroup.Done()
	}()
	go func() {
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inC, ChOut: inD}
		program.Run()
		waitgroup.Done()
	}()
	go func() {
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inD, ChOut: inE}
		program.Run()
		waitgroup.Done()
	}()
	go func() {
		program := intcode.Program{Memory: append(baseProgram[:0:0], baseProgram...), ChIn: inE, ChOut: feedBack}
		program.Run()
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
