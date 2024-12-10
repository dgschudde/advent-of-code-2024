package main

import (
	"advent-of-code-2024/common"
	"fmt"
)

func main() {
	var input *[]string
	input = common.ReadInput("./input/test-input.txt")

	convertedInput := convertInput(*input)
	occurrences := findOccurrences(convertedInput)

	fmt.Printf("The word MAS forms a cross %d times.\r\n", occurrences)
}

func convertInput(input []string) [][]rune {
	convertedInput := make([][]rune, len(input))
	for i, row := range input {
		convertedInput[i] = []rune(row)
	}
	return convertedInput
}

func findOccurrences(grid [][]rune) int {
	var count int
	word := []rune("MAS")

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if checkWord(grid, i, j, word) {
				count++
			}
		}
	}

	return count
}

func checkWord(grid [][]rune, row, col int, word []rune) bool {
	if row-1 >= 0 && row+1 < len(grid) && col-1 >= 0 && col+1 < len(grid[0]) {
		if grid[row-1][col] == word[0] && grid[row][col-1] == word[0] && grid[row][col] == word[1] && grid[row][col+1] == word[2] && grid[row+1][col] == word[2] {
			return true
		}
	}
	return false
}
