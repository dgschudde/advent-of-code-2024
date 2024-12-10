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

	fmt.Printf("The word XMAS occurred %d times.\r\n", len(occurrences))
}

func convertInput(input []string) [][]rune {
	convertedInput := make([][]rune, len(input))
	for i, row := range input {
		convertedInput[i] = []rune(row)
	}
	return convertedInput
}

type Coordinate struct {
	Row int
	Col int
}

func findOccurrences(grid [][]rune) []Coordinate {
	var occurrences []Coordinate
	word := []rune("MAS")

	directions := [][2]int{
		{1, 1},   // diagonal down-right
		{-1, -1}, // diagonal up-left
		{-1, 1},  // diagonal up-right
		{1, -1},  // diagonal down-left
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if checkWord(grid, i, j, dir, word) {
					occurrences = append(occurrences, Coordinate{i, j})
				}
			}
		}
	}

	return occurrences
}

func checkWord(grid [][]rune, row, col int, dir [2]int, word []rune) bool {
	for k := 0; k < len(word); k++ {
		newRow := row + k*dir[0]
		newCol := col + k*dir[1]
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || (k == 1 && grid[newRow][newCol] != word[k]) || (k != 1 && grid[newRow][newCol] != word[k]) {
			return false
		}
	}
	return true
}
