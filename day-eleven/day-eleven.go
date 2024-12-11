package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"strconv"
	"strings"
)

type Rule int

const (
	Swap Rule = iota
	Multiply
	Split
)

var RuleName = map[Rule]string{
	Swap:     "Swap",
	Multiply: "Multiply",
	Split:    "Split",
}

func (ss Rule) String() string {
	return RuleName[ss]
}

func main() {
	var input *[]string
	input = common.ReadInput("./input/input.txt")

	var stones []uint64

	stones = convertInput(*input)
	refStones := &stones
	amount := CalculateStones(refStones)

	fmt.Printf("Stones found: %d", amount)
}

func convertInput(input []string) []uint64 {
	var numbers []uint64

	s := strings.Split(input[0], " ")
	for _, v := range s {
		number, _ := strconv.ParseUint(v, 10, 64)
		numbers = append(numbers, number)
	}
	return numbers
}

func CalculateStones(stones *[]uint64) int {
	fmt.Printf("Starting with %d stones\r\n", len(*stones))

	for blink := 1; blink <= 75; blink++ {
		var delta = 0

		fmt.Printf("Processing blink: %d\r\n", blink)
		for i := 0; i < len(*stones); i++ {
			refStones := *stones
			i += delta
			if i+delta > len(*stones) {
				break
			}
			delta = 0

			var rule = CheckRuleApplies(refStones[i])
			switch rule {
			case Swap:
				SwapValue(stones, i)
			case Multiply:
				MultiplyValue(stones, i)
			case Split:
				stones, delta = SplitValues(stones, i)
			default:
				panic("Rule doesn't exist")
			}
		}
	}

	return len(*stones)
}

func CheckRuleApplies(stone uint64) Rule {
	if stone == 0 {
		return Swap
	} else if IsEven(stone) {
		return Split
	}
	return Multiply
}

func IsEven(stone uint64) bool {
	stoneString := strconv.FormatUint(stone, 10)
	return len(stoneString)%2 == 0
}

func SwapValue(numbers *[]uint64, index int) *[]uint64 {
	refNumbers := *numbers
	refNumbers[index] = 1
	return &refNumbers
}

func MultiplyValue(numbers *[]uint64, index int) *[]uint64 {
	refNumbers := *numbers
	v := refNumbers[index] * 2024
	refNumbers[index] = v
	return &refNumbers
}

func SplitValues(numbers *[]uint64, index int) (*[]uint64, int) {
	refNumbers := *numbers

	numberString := strconv.FormatUint(refNumbers[index], 10)

	half := len(numberString) / 2
	leftString := numberString[:half]
	rightString := numberString[half:]

	left, _ := strconv.ParseUint(leftString, 10, 64)
	right, _ := strconv.ParseUint(rightString, 10, 64)

	refNumbers = remove(&refNumbers, index)
	refNumbers = insert(&refNumbers, index, left)
	refNumbers = insert(&refNumbers, index+1, right)

	return &refNumbers, 1
}

func insert(a *[]uint64, index int, value uint64) []uint64 {
	refA := *a
	if len(refA) == index {
		refA = append(refA, value)
		return refA
	}
	refA = append(refA[:index+1], refA[index:]...)
	refA[index] = value
	return refA
}

func remove(slice *[]uint64, s int) []uint64 {
	refSlice := *slice
	refSlice = append(refSlice[:s], refSlice[s+1:]...)
	return refSlice
}
