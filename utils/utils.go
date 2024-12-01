package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(test bool, scanFunc func(scanner *bufio.Scanner)) {
	fileName := "input.txt"
	if test {
		fileName = "input_test.txt"
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanFunc(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
