package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var lefts []int
var rights []int

func main() {
	var input *[]string

	input = common.ReadInput()
	ConvertInput(input)
	var totalDistance int
	frequencyMap := checkFrequency(lefts, rights)
	for _, value := range lefts {
		totalDistance += value * frequencyMap[value]
	}

	fmt.Printf("Total distance is %d\n", totalDistance)
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
