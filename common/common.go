package common

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(fileName string) *[]string {
	var input = make([]string, 0)

	// Read the input from file
	inputFile, err := os.Open(fileName)

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
