package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var input *[]string
	var inputToString string

	input = common.ReadInput("./input/input.txt")

	for _, currentLine := range *input {
		inputToString += currentLine
	}

	instructions := matchInstructions(inputToString)
	total := calculateTotal(instructions)

	fmt.Println(total)
}

func calculateTotal(instructions []string) int64 {
	var regex = regexp.MustCompile(`\d+`)

	var total int64
	var multiplyEnabled = true

	for _, instruction := range instructions {
		if instruction == `don't()` {
			multiplyEnabled = false
			continue
		}
		if instruction == `do()` {
			multiplyEnabled = true
			continue
		}

		if multiplyEnabled {
			numbers := regex.FindAllString(instruction, -1)
			left, _ := strconv.ParseInt(numbers[0], 10, 64)
			right, _ := strconv.ParseInt(numbers[1], 10, 64)
			total = total + left*right
		}
	}
	return total
}

func matchInstructions(input string) []string {
	regex := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	matches := regex.FindAllString(input, -1)
	return matches
}
