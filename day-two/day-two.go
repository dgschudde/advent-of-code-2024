package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input *[]string

	input = ReadInput()

	var totalSafeReports int

	for _, currentLine := range *input {
		series := convertStringArrayToIntArray(strings.Split(currentLine, " "))
		result := checkSeries(series)
		if result == "safe" {
			totalSafeReports++
		}
	}

	fmt.Printf("Total safe reports: %d\r\n", totalSafeReports)
}

func convertStringArrayToIntArray(strArray []string) []int {
	intArray := make([]int, len(strArray))
	for i, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		intArray[i] = num
	}
	return intArray
}

func checkSeries(series []int) string {
	increasing := false
	decreasing := false

	for i := 0; i < len(series)-1; i++ {
		diff := series[i+1] - series[i]
		if diff > 0 {
			increasing = true
		} else if diff < 0 {
			decreasing = true
		} else if diff == 0 {
			return "unsafe"
		}

		if diff > 3 || diff < -3 || (increasing && decreasing) {
			return "unsafe"
		}
	}
	return "safe"
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
