/*
 Saw a solution like this on reddit and loved it.
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"unicode"
)

type Number struct {
	X     int
	Y     int
	Value string
}

type Symbol struct {
	X               int
	Y               int
	Value           string
	AdjacentNumbers []int
}

func main() {
	lines, _ := utils.GetPuzzleInputAsLines("https://adventofcode.com/2023/day/3/input")

	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)
	for y, line := range lines {
		line = line + "."
		n := Number{}
		s := Symbol{}
		for x, char := range line {
			if !unicode.IsDigit(char) && char != '.' {
				s.X = x
				s.Y = y
				s.Value = string(char)
				symbols = append(symbols, s)
				s = Symbol{}
			}
			if !unicode.IsDigit(char) && n.Value != "" {
				n.X = x - len(n.Value)
				n.Y = y
				numbers = append(numbers, n)
				n = Number{}
			}
			if unicode.IsDigit(char) {
				n.Value += string(char)
			}
		}
	}

	// Part one
	sum := 0
	for _, n := range numbers {
		for i := range symbols {
			minX := n.X - 1
			minY := n.Y - 1
			maxX := n.X + len(n.Value)
			maxY := n.Y + 1
			s := &symbols[i]
			if minX <= s.X && s.X <= maxX && minY <= s.Y && s.Y <= maxY {
				v, _ := strconv.Atoi(n.Value)
				s.AdjacentNumbers = append(s.AdjacentNumbers, v)
				sum += v
			}
		}
	}

	fmt.Println("The sum of all the part numbers is:")
	fmt.Println(sum)

	// Part two
	sum = 0
	for _, s := range symbols {
		if s.Value == "*" {
			if len(s.AdjacentNumbers) == 2 {
				sum += s.AdjacentNumbers[0] * s.AdjacentNumbers[1]
			}
		}
	}
	fmt.Println("The sum of all the products of the gear ratios is equal to:")
	fmt.Println(sum)

}
