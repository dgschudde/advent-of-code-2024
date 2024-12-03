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
		} else {
			unsafeSeries := convertStringArrayToIntArray(strings.Split(currentLine, " "))
			if checkUnsafeSeries(unsafeSeries) {
				totalSafeReports++
			}
		}
	}

	fmt.Printf("Total safe reports: %d\r\n", totalSafeReports)
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

func checkUnsafeSeries(series []int) bool {
	var copiedSeries []int = make([]int, len(series))

	for i := 0; i < len(series); i++ {
		copy(copiedSeries, series)
		removedSeries := remove(copiedSeries, i)
		result := checkSeries(removedSeries)
		if result == "safe" {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
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
		var diff = series[i+1] - series[i]
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
