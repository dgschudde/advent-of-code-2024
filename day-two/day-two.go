package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input *[]string

	input = common.ReadInput("./input/input.txt")

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
