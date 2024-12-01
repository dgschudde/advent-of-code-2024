package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lefts []int
var rights []int

func main() {
	var input *[]string

	input = ReadInput()
	ConvertInput(input)
	distance := CalculateDistance()

	fmt.Println(distance)
}

func ReadInput() *[]string {
	var input = make([]string, 0)

	// Read the input from file
	inputFile, err := os.Open("./input/test-input.txt")

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

func ConvertInput(input *[]string) {
	for _, currentLine := range *input {
		left, _ := strconv.ParseInt(strings.Split(currentLine, "   ")[0], 10, 64)
		right, _ := strconv.ParseInt(strings.Split(currentLine, "   ")[1], 10, 64)
		lefts = append(lefts, int(left))
		rights = append(rights, int(right))
	}

	sort.Ints(lefts)
	sort.Ints(rights)
}

func CalculateDistance() int64 {
	var totalDistance int64
	var distance int

	for i := 0; i < len(lefts); i++ {
		left := lefts[i]
		right := rights[i]
		if left >= right {
			distance = left - right
		} else {
			distance = right - left
		}

		totalDistance += int64(distance)
	}
	return totalDistance
}
