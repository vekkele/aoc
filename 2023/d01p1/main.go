package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("d01p1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := getCalibrationValue(scanner.Text())
		if err != nil {
			log.Printf("failed to parse \"%s\". Error: %v\n", line, err)
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", sum)
}

func getCalibrationValue(line string) (int, error) {
	var left, right string

	for i, j := 0, len(line)-1; i < len(line); i, j = i+1, j-1 {
		if left != "" && right != "" {
			break
		}

		leftChar := rune(line[i])
		rightChar := rune(line[j])

		isLeftNum := unicode.IsDigit(leftChar)
		isRightNum := unicode.IsDigit(rightChar)

		if isLeftNum && left == "" {
			left = string(leftChar)
		}

		if isRightNum && right == "" {
			right = string(rightChar)
		}
	}

	return strconv.Atoi(left + right)
}
