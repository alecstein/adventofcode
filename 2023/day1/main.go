package main

/*
To check for the last word-number in the string, I reversed the string
and searched for the earliest reversed number. That was the only non-trivial
part of this problem.
*/

import (
	"adventofcode/utils"
	"fmt"
	"strings"
	"unicode"
)

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Returns (index, value) where index is the index of the first digit
// and value is its value
func firstDigit(s string) (int, int) {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return i, int(r - 48)
		}
	}
	return -1, -1
}

// Returns (index, value) where index is the index of the last digit
// and value is its value. Index is measured from the end of the string
func lastDigit(s string) (int, int) {
	return firstDigit(utils.Reverse(s))
}

// Returns (index, value) where index is the index of the first (word) number
// and value is its value
func firstNumber(m map[string]int, s string) (int, int) {
	wIndex, wValue := -1, -1

	for k, v := range m {
		i := strings.Index(s, k)
		if i == -1 {
			continue
		}
		if wIndex == -1 || i < wIndex {
			wIndex = i
			wValue = v
		}
	}
	return wIndex, wValue
}

// Returns (index, value) where index is the index of the first digit or number
// and value is its value
func firstDigitOrNumber(m map[string]int, s string) (int, int) {
	wIndex, wValue := firstNumber(m, s)
	dIndex, dValue := firstDigit(s)
	if wIndex == -1 {
		return dIndex, dValue
	}
	if dIndex == -1 {
		return wIndex, wValue
	}
	if wIndex < dIndex {
		return wIndex, wValue
	}
	return dIndex, dValue
}

// Returns (index, value) where index is the index of the last digit or number
// and value is its value. Index is measured from the end of the string
func lastDigitOrNumber(m map[string]int, s string) (int, int) {
	revMap := make(map[string]int)
	for k, v := range nums {
		revMap[utils.Reverse(k)] = v
	}
	return firstDigitOrNumber(revMap, utils.Reverse(s))
}

func main() {
	lines, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/1/input")

	// Part one
	s := 0
	for _, line := range lines {
		_, first := firstDigit(line)
		_, last := lastDigit(line)
		twoDigitNumber := first*10 + last
		s += twoDigitNumber
	}
	fmt.Printf("The sum of the first and last digits in each string is:\n%d\n", s)

	// Part two
	s = 0
	for _, line := range lines {
		_, first := firstDigitOrNumber(nums, line)
		_, last := lastDigitOrNumber(nums, line)
		twoDigitNumber := first*10 + last
		s += twoDigitNumber
	}
	fmt.Printf("The sum of the first and last **number of any kind** in each string is:\n%d\n", s)
}
