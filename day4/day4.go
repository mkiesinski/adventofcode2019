package main

import (
	"fmt"
	"strconv"
)

func verify(code int) bool {
	repeated := false
	codeString := strconv.Itoa(code)

	for i := 0; i < len(codeString)-1; i++ {
		if codeString[i] > codeString[i+1] {
			return false
		}
		if codeString[i] == codeString[i+1] {
			repeated = true
		}
	}

	return repeated
}

func verify2(code int) bool {
	codeString := strconv.Itoa(code)
	sequence := 1
	hasDouble := false

	if codeString[0] > codeString[1] {
		return false
	}

	for i := 0; i < len(codeString)-1; i++ {
		if codeString[i] > codeString[i+1] {
			return false
		}
		if codeString[i] == codeString[i+1] {
			sequence++
		} else if sequence != 1 {
			if sequence == 2 {
				hasDouble = true
			}
			sequence = 1
		}
	}

	if sequence != 1 && sequence == 2 {
		hasDouble = true
	}

	return hasDouble
}

func findAllCodes(start int, end int) []int {
	results := make([]int, 0)

	for i := start; i <= end; i++ {
		if verify(i) {
			results = append(results, i)
		}
	}

	return results
}

func findAllCodesRefined(start int, end int) []int {
	results := make([]int, 0)

	for i := start; i <= end; i++ {
		if verify2(i) {
			results = append(results, i)
		}
	}

	return results
}

func main() {
	start := 124075
	end := 580769

	fmt.Println("=== Part 1 ===")
	fmt.Printf("All codes meeting the requirement between %d and %d : %d\n", start, end, len(findAllCodes(start, end)))
	fmt.Println("=== Part 2 ===")
	fmt.Printf("All codes meeting the requirement between %d and %d : %d\n", start, end, len(findAllCodesRefined(start, end)))
}
