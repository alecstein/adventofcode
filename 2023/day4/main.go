/*
I learned a bit about regex and string indexing. It turns
out the FieldsFunc is a much nicer way of splitting strings,
and more efficient than regex.

I made some effor to solve this in a single pass and without structs.
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	// input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/4/input")
	// Optional local import
	input, _ := utils.GetPuzzleInputFromFile("input.txt")

	// input = []string{
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	// }

	// Used for Part two
	copiesArr := make([]int, len(input))
	for i := range input {
		copiesArr[i] = 1
	}

	// Part one
	sum := 0

	for i, line := range input {
		// FieldsFunc is a great discovery!
		// https://pkg.go.dev/strings#FieldsFunc
		splitLine := strings.FieldsFunc(
			line,
			func(r rune) bool {
				return r == ':' || r == '|'
			},
		)

		winningNumbersString := splitLine[1]
		winningNumbers := strings.FieldsFunc(
			winningNumbersString,
			func(r rune) bool {
				return r == ' '
			},
		)

		numbersString := splitLine[2]
		numbers := strings.FieldsFunc(
			numbersString,
			func(r rune) bool {
				return r == ' '
			},
		)

		// Part one
		points := 0
		for _, n := range numbers {
			for _, w := range winningNumbers {
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

		// Part two (separated clarity)
		score := 0
		for _, n := range numbers {
			for _, w := range winningNumbers {
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

	fmt.Println("The sum of all the points is:")
	fmt.Println(sum)

	sum = 0
	for _, copies := range copiesArr {
		sum += copies
	}

	fmt.Println("The number of scratchcards you end up with is:")
	fmt.Println(sum)

}
