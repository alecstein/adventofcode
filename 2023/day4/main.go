/*
I learned a bit about regex and string indexing.
My initial solution used lots of structs, when it turned out
arrays worked just fine.

I didn't make much effort to super-optimize this, but I did decide
to try solving it in a single pass, which is why Part one and Part two
are interleaved.
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

type Card struct {
	WinningNumbers []string
	Numbers        []string
	CurrentCopies  int
}

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/4/input")
	// Optional local import
	// input, _ = utils.GetPuzzleInputFromFile("input.txt")

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

	// Used for Part two
	copiesArr := make([]int, len(input))
	for i := range input {
		copiesArr[i] = 1
	}

	// Part one
	sum := 0
	for i, line := range input {
		winningNumbersString := line[colIdx+1 : pipIdx-1]
		numPat := regexp.MustCompile(`(\d+)`)
		winningNumberMatches := numPat.FindAllString(winningNumbersString, -1)
		numbersString := line[pipIdx+1:]
		numbersMatches := numPat.FindAllString(numbersString, -1)

		// Part one
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

		// Part two (separated clarity)
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

	fmt.Println("The sum of all the points is:")
	fmt.Println(sum)

	sum = 0
	for _, copies := range copiesArr {
		sum += copies
	}

	fmt.Println("The number of scratchcards you end up with is:")
	fmt.Println(sum)

}
