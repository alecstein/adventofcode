package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func isPartNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) && string(c) != "." {
			return true
		}
	}
	return false
}

func main() {
	lines, _ := utils.GetPuzzleInputAsLines("https://adventofcode.com/2023/day/3/input")

	// Make a line of N dots, where N is equal to the length of one of the lines
	// of lines
	dummyLine := ""
	for i := 0; i < len(lines[0]); i++ {
		dummyLine += "."
	}
	lines = append(lines, dummyLine)
	lines = append([]string{dummyLine}, lines...)

	// Add padding to the lines
	for i, line := range lines {
		lines[i] = "." + line + "."
	}

	// Part one
	s := 0
	for i, line := range lines[1 : len(lines)-1] {
		re := regexp.MustCompile(`(\d+)`)
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			startIdx := match[0]
			endIdx := match[1]
			leftChar := string(line[startIdx-1])
			rightChar := string(line[endIdx])
			topChars := string(lines[i][startIdx-1 : endIdx+1])
			bottomChars := string(lines[i+2][startIdx-1 : endIdx+1])
			allChars := leftChar + rightChar + topChars + bottomChars

			if isPartNumber(allChars) {
				k, err := strconv.Atoi(line[startIdx:endIdx])
				if err != nil {
					panic(err)
				}

				s += k
			}
		}
	}
	fmt.Println("The sum of all the part numbers is:")
	fmt.Println(s)
}
