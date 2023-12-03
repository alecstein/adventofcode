package main

/*
For part one, I solved it by collecting the border of each number and
checking if that border was all periods (.) or not. (For edges, I added
a padding of periods to avoid out-of-bounds errors.)

For part two, notice that any number adjacent to a star is by definition a
part number. I looped over each number and figured out whether it was touching a
star. If it was, I made a map
	star -> []int
where the slice was the numbers that touched that star. Then I looped over the stars
to see which ones had slices of length two, i.e. they were touching two numbers.
*/

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// A number is a part number if at least one of its border characters
// is not a period (or a digit)
func isPartNumber(borderString string) bool {
	for _, c := range borderString {
		if !unicode.IsDigit(c) && string(c) != "." {
			return true
		}
	}
	return false
}

type Star struct {
	Row, Col int
}

func main() {
	lines, _ := utils.GetPuzzleInputAsLines("https://adventofcode.com/2023/day/3/input")

	/*
		To avoid out-of-bounds errors, we add a padding of periods to the top, bottom,
		and sides.
	*/
	paddingLine := strings.Repeat(".", len(lines[0]))
	lines = append(lines, paddingLine)
	lines = append([]string{paddingLine}, lines...)
	for i, line := range lines {
		lines[i] = "." + line + "."
	}

	// Part one
	s := 0
	for i, line := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		// Get the start and end indices of the numbers on this line
		re := regexp.MustCompile(`(\d+)`)
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			// For each number found, get the characters in the bounding box around it
			begIdx := match[0]
			endIdx := match[1]

			charLeft := string(line[begIdx-1])
			charRight := string(line[endIdx])
			charsTop := string(lines[i-1][begIdx-1 : endIdx+1])
			charsBottom := string(lines[i+1][begIdx-1 : endIdx+1])

			charsBorder := charLeft + charRight + charsTop + charsBottom

			if isPartNumber(charsBorder) {
				num, _ := strconv.Atoi(line[begIdx:endIdx])
				s += num
			}
		}
	}
	fmt.Println("The sum of all the part numbers is:")
	fmt.Println(s)

	// Part two
	/*
		We're going to do this the following way, which I think is pretty clean.
		We go through all the numbers on each line and record the locations of the stars
		they touch. Then we put those star locations in a map:
			star -> []int
		Every time we find a number that touches a star, we add that number to the slice
		associated with that star.
	*/
	stars := make(map[Star][]int, 0)
	for i, line := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}
		re := regexp.MustCompile(`(\d+)`)
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			begIdx, endIdx := match[0], match[1]
			num, _ := strconv.Atoi(line[begIdx:endIdx])

			charsAbove := lines[i-1][begIdx-1 : endIdx+1]
			for j, c := range charsAbove {
				if string(c) == "*" {
					star := Star{Row: i - 1, Col: begIdx - 1 + j}
					stars[star] = append(stars[star], num)
				}
			}

			charsBelow := lines[i+1][begIdx-1 : endIdx+1]
			for j, c := range charsBelow {
				if string(c) == "*" {
					star := Star{Row: i + 1, Col: begIdx - 1 + j}
					stars[star] = append(stars[star], num)
				}
			}

			charLeft := string(line[begIdx-1])
			if charLeft == "*" {
				star := Star{Row: i, Col: begIdx - 1}
				stars[star] = append(stars[star], num)
			}

			charRight := string(line[endIdx])
			if charRight == "*" {
				star := Star{Row: i, Col: endIdx}
				stars[star] = append(stars[star], num)
			}
		}
	}

	s = 0
	for _, v := range stars {
		if len(v) == 2 {
			s += v[0] * v[1]
		}
	}

	fmt.Println("The sum of all the products of the gear ratios is equal to:")
	fmt.Println(s)
}
