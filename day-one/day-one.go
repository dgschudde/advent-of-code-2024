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
	var totalDistance int
	frequencyMap := checkFrequency(lefts, rights)
	for _, value := range lefts {
		totalDistance += value * frequencyMap[value]
	}

	fmt.Printf("Total distance is %d\n", totalDistance)
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

func checkFrequency(array1, array2 []int) map[int]int {
	frequencyMap := make(map[int]int)
	for _, value := range array1 {
		frequencyMap[value] = 0
		for _, v := range array2 {
			if value == v {
				frequencyMap[value]++
			}
		}
	}
	return frequencyMap
}
