package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Advent of Code - Day 1\n")

	fmt.Println("Advent of Code - Part 1\n")

	sampleResult := solveDayOnePuzzle("day-1/sample-1.txt")
	fmt.Printf("Sample Total is %v\n", sampleResult)

	puzzleResult := solveDayOnePuzzle("day-1/puzzle.txt")
	fmt.Printf("Puzzle Total is %v\n", puzzleResult)

	fmt.Println()

	fmt.Println("Advent of Code - Part 2\n")

	sampleAdvancedResult := solveDayOneAdvancedPuzzle("day-1/sample-2.txt")
	fmt.Printf("Advanced Sample Total is %v\n", sampleAdvancedResult)

	puzzleAdvancedResult := solveDayOneAdvancedPuzzle("day-1/puzzle.txt")
	fmt.Printf("Advanced Puzzle Total is %v\n", puzzleAdvancedResult)
}

func solveDayOnePuzzle(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var total int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		total = total + lineTotal(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func solveDayOneAdvancedPuzzle(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var total int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		textConvert := ""

		numbers := make(map[string]int)
		numbers["one"] = 1
		numbers["two"] = 2
		numbers["three"] = 3
		numbers["four"] = 4
		numbers["five"] = 5
		numbers["six"] = 6
		numbers["seven"] = 7
		numbers["eight"] = 8
		numbers["nine"] = 9

		currentWord := text
		for _, char := range text {
			if unicode.IsNumber(char) {
				textConvert += string(char)
				currentWord = trimLeftChar(currentWord)
			} else {
				for key, value := range numbers {
					if strings.HasPrefix(currentWord, key) {
						textConvert += fmt.Sprintf("%v", value)
						break
					}
				}
				currentWord = trimLeftChar(currentWord)
			}
		}

		lineTotal := lineTotal(textConvert)
		total = total + lineTotal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return total
}

func lineTotal(textConvert string) int64 {
	first := ""
	second := ""

	for _, char := range textConvert {
		if unicode.IsNumber(char) {
			if first == "" {
				first = string(char)
			} else {
				second = string(char)
			}
		}
	}
	if first == "" {
		first = "0"
	}
	if second == "" {
		second = first
	}
	strVal := fmt.Sprintf("%v%v", first, second)
	value, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
