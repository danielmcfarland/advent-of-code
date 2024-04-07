package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 2\n")
	// 12 red cubes, 13 green cubes, 14 blue cubes

	fmt.Println("Advent of Code - Part 1\n")

	solveDayTwoPuzzle("day-2/sample-1.txt")
	solveDayTwoPuzzle("day-2/puzzle.txt")

	fmt.Println()
	fmt.Println("Advent of Code - Part 2\n")
}

func solveDayTwoPuzzle(filename string) {
	colors := make(map[string]int)
	colors["red"] = 12
	colors["green"] = 13
	colors["blue"] = 14

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	gameNumber := 1
	gameIdCount := 0
	for scanner.Scan() {
		gamePossible := true
		text := scanner.Text()
		res1 := strings.Split(text, ": ")
		res2 := strings.Split(res1[1], "; ")
		for _, v := range res2 {
			res3 := strings.Split(v, ", ")
			for _, v2 := range res3 {

				res4 := strings.Split(v2, " ")

				for key, value := range colors {
					if strings.HasSuffix(res4[1], key) {
						gameValue, err := strconv.ParseInt(res4[0], 10, 64)
						if err != nil {
							log.Fatal(err)
						}
						if gameValue > int64(value) {
							gamePossible = false
							break
						}
					}
				}
			}
		}

		if gamePossible {
			gameIdCount += gameNumber
		}

		gameNumber++
	}

	fmt.Println("Sum of possible game IDs: ", gameIdCount)
}
