package main

import (
	"advent-of-code-2024/common"
	"fmt"
	"strings"
)

type Computer struct {
	name         string
	destinations []Computer
}

func main() {
	var input *[]string
	input = common.ReadInput("./input/test-input.txt")

	network := convertInput(*input)

	fmt.Println(len(*network))
}

func convertInput(input []string) *[]Computer {
	network := make([]Computer, 0)

	for i := 0; i < len(input); i++ {
		values := strings.Split(input[i], "-")

		for j := 0; j < len(values); j++ {
			computer := findComputer(network, values[j])

			if computer == nil {
				if len(network) > 0 {
					fmt.Printf("Adding %s -> %s\r\n", network[len(network)-1].name, values[j])
					network = append(network, Computer{values[j], nil})
					target := findComputer(network, values[j])
					if target != nil {
						fmt.Printf("Adding destination %s to %s\r\n", target.name, network[i].name)
						network[i].destinations = append(network[i].destinations, *target)
					}
				} else {
					network = append(network, Computer{values[j], nil})
				}

			}

			fmt.Println(computer)
		}

		fmt.Printf("Addding %s -> %s\r\n", values[0], values[1])
	}

	return &network
}

func findComputer(network []Computer, computerName string) *Computer {
	fmt.Printf("Search computer %s\r\n", computerName)
	for i := 0; i < len(network); i++ {
		if network[i].name == computerName {
			fmt.Printf("Found computer %s\r\n", computerName)
			return &(network)[i]
		}
	}
	return nil
}
