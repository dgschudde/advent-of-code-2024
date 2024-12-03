package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var input *[]string
	var inputToString string

	input = ReadInput()

	for _, currentLine := range *input {
		inputToString += currentLine
	}

	instructions := matchInstructions(inputToString)
	total := calculateTotal(instructions)

	fmt.Println(total)
	fmt.Printf("Read %d lines of input", len(*input))
}

func calculateTotal(instructions []string) int64 {
	var regex = regexp.MustCompile(`\d+`)

	var total int64

	for _, instruction := range instructions {
		numbers := regex.FindAllString(instruction, -1)
		left, _ := strconv.ParseInt(numbers[0], 10, 64)
		right, _ := strconv.ParseInt(numbers[1], 10, 64)
		total = total + left*right
	}
	return total
}

func matchInstructions(input string) []string {
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := regex.FindAllString(input, -1)
	return matches
}

func ReadInput() *[]string {
	var input = make([]string, 0)

	// Read the input from file
	inputFile, err := os.Open("./input/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return &input
}
