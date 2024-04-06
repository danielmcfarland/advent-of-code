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
	fmt.Println("Advent of Code - Day 1\n")

	fmt.Println("Advent of Code - Part 1\n")

	sampleResult := solvePuzzle("day-1/sample-1.txt")
	puzzleResult := solvePuzzle("day-1/puzzle.txt")

	fmt.Printf("Sample Total is %v\n", sampleResult)
	fmt.Printf("Puzzle Total is %v\n", puzzleResult)

	fmt.Println()

	fmt.Println("Advent of Code - Part 2\n")
}

func solvePuzzle(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var total int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		first := ""
		second := ""

		for _, char := range text {
			if unicode.IsNumber(char) {
				if first == "" {
					first = string(char)
				} else {
					second = string(char)
				}
			}
		}
		if second == "" {
			second = first
		}
		strVal := fmt.Sprintf("%v%v", first, second)
		value, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total = total + value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}
