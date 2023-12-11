package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/8/input")

	var instructions string
	network := make(map[string][2]string, 0)
	for i, line := range input {
		if i == 0 { // handle first line separately
			instructions = line
		} else if line != "" {
			fields := strings.Fields(line)
			k := fields[0]
			v1 := fields[2]
			v1 = v1[1 : len(v1)-1]
			v2 := fields[3]
			v2 = v2[:len(v2)-1]
			network[k] = [2]string{v1, v2}
		}
	}

	counts := 0
	loc := "AAA"
	i := 0
	for loc != "ZZZ" {
		counts++
		j := i % len(instructions)
		v := 0
		if instructions[j] == 'R' {
			v = 1
		}
		loc = network[loc][v]
		i++
	}

	fmt.Println(counts)
}
