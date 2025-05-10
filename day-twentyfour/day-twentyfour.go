package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := common.ReadInput("input/input.txt")
	wireValues, gateConnections := parseWireValues(*input)

	initialConnections := make([]string, 0)
	remainingConnections := make([]string, 0)
	for _, connection := range gateConnections {
		parts := strings.Split(connection, " ")
		if strings.HasPrefix(parts[0], "x") || strings.HasPrefix(parts[0], "y") || strings.HasPrefix(parts[2], "x") || strings.HasPrefix(parts[2], "y") {
			initialConnections = append(initialConnections, connection)
		} else {
			remainingConnections = append(remainingConnections, connection)
		}
	}

	for _, connection := range initialConnections {
		parts := strings.Split(connection, " ")
		input1 := wireValues[parts[0]]
		input2 := wireValues[parts[2]]
		output := parts[4]

		switch parts[1] {
		case "AND":
			wireValues[output] = andGate(input1, input2)
		case "OR":
			wireValues[output] = orGate(input1, input2)
		case "XOR":
			wireValues[output] = xorGate(input1, input2)
		}
	}

	// Process remaining gate connections sorted by output in alphabetical order
	sort.Slice(remainingConnections, func(i, j int) bool {
		partsI := strings.Split(remainingConnections[i], " ")
		partsJ := strings.Split(remainingConnections[j], " ")
		return partsI[4] < partsJ[4]
	})

	processedConnections := make(map[string]bool)
	for len(remainingConnections) > 0 {
		unprocessedConnections := make([]string, 0)
		for _, connection := range remainingConnections {
			parts := strings.Split(connection, " ")
			input1, input1Exists := wireValues[parts[0]]
			input2, input2Exists := wireValues[parts[2]]
			output := parts[4]

			if input1Exists && input2Exists {
				switch parts[1] {
				case "AND":
					wireValues[output] = andGate(input1, input2)
				case "OR":
					wireValues[output] = orGate(input1, input2)
				case "XOR":
					wireValues[output] = xorGate(input1, input2)
				}
				processedConnections[connection] = true
			} else {
				unprocessedConnections = append(unprocessedConnections, connection)
			}
		}
		if len(unprocessedConnections) == len(remainingConnections) {
			break
		}
		remainingConnections = unprocessedConnections
	}

	outputBits := ""
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if value, exists := wireValues[wire]; exists {
			outputBits = strconv.Itoa(value) + outputBits
		} else {
			break
		}
	}

	outputDecimal, _ := strconv.ParseInt(outputBits, 2, 64)
	fmt.Println(outputDecimal)
}

func parseWireValues(input []string) (map[string]int, []string) {
	wireValues := make(map[string]int)
	gateConnections := make([]string, 0)

	for _, line := range input {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) == 2 {
			wireValues[parts[0]], _ = strconv.Atoi(parts[1])
		} else {
			gateConnections = append(gateConnections, line)
		}
	}
	return wireValues, gateConnections
}

func andGate(input1, input2 int) int {
	if input1 == 1 && input2 == 1 {
		return 1
	}
	return 0
}

func orGate(input1, input2 int) int {
	if input1 == 0 && input2 == 0 {
		return 0
	}
	return 1
}

func xorGate(input1, input2 int) int {
	if input1 != input2 {
		return 1
	}
	return 0
}
