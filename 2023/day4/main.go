/*
I learned a bit about regex and string indexing.
My initial solution used lots of structs, when it turned out
arrays worked just fine.
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input, _ := utils.GetPuzzleInputAsLines("https://adventofcode.com/2023/day/4/input")

	// input = []string{
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	// }

	colIdx := strings.Index(input[0], ":")
	pipIdx := strings.Index(input[0], "|")

	// Part one
	sum := 0
	for _, line := range input {
		winningNumbersString := line[colIdx+1 : pipIdx-1]
		numPat := regexp.MustCompile(`(\d+)`)
		winningNumberMatches := numPat.FindAllString(winningNumbersString, -1)
		numbersString := line[pipIdx+1:]
		numbersMatches := numPat.FindAllString(numbersString, -1)
		points := 0
		for _, n := range numbersMatches {
			for _, w := range winningNumberMatches {
				if n == w {
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		sum += points
	}

	fmt.Println("The sum of all the points is:")
	fmt.Println(sum)

	// Part two
	copiesArr := make([]int, len(input))
	for i := range input {
		copiesArr[i] = 1
	}
	for i, line := range input {
		winningNumbersString := line[colIdx+1 : pipIdx-1]
		numPat := regexp.MustCompile(`(\d+)`)
		winningNumberMatches := numPat.FindAllString(winningNumbersString, -1)
		numbersString := line[pipIdx+1:]
		numbersMatches := numPat.FindAllString(numbersString, -1)
		score := 0
		for _, n := range numbersMatches {
			for _, w := range winningNumberMatches {
				if n == w {
					score++
				}
			}
		}
		for j := 1; j <= score; j++ {
			if i+j < len(input) {
				copiesArr[i+j] += copiesArr[i]
			}
		}
	}
	sum = 0
	for _, copies := range copiesArr {
		sum += copies
	}

	fmt.Println("The number of scratchcards you end up with is:")
	fmt.Println(sum)
}
