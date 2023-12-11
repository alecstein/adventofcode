/* Did not like this puzzle.

Coding up the naive algorithm was fast, maybe 15 minutes for both parts, and was fun.

But I learned (from reading on reddit) that you have to find the LCM. This isn't really
inferrable from the problem statement. You kind of have to guess.
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/8/input")
	// input = []string{
	// 	"LR",
	// 	"",
	// 	"11A = (11B, XXX)",
	// 	"11B = (XXX, 11Z)",
	// 	"11Z = (11B, XXX)",
	// 	"22A = (22B, XXX)",
	// 	"22B = (22C, 22C)",
	// 	"22C = (22Z, 22Z)",
	// 	"22Z = (22B, 22B)",
	// 	"XXX = (XXX, XXX)",
	// }

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
		// Circular array
		j := i % len(instructions)
		v := 0
		if instructions[j] == 'R' {
			v = 1
		}
		loc = network[loc][v]
		i++
	}

	fmt.Println(counts)

	locs := make([]string, 0)
	for k := range network {
		if k[2] == 'A' {
			locs = append(locs, k)
		}
	}

	countsArr := make([]int, 0)

	for _, loc := range locs {
		counts := 0
		i := 0
		for loc[2] != 'Z' {
			counts++
			j := i % len(instructions)
			v := 0
			if instructions[j] == 'R' {
				v = 1
			}
			loc = network[loc][v]
			i++
		}
		countsArr = append(countsArr, counts)
	}

	lcm := 1
	for _, c := range countsArr {
		lcm = utils.LCM(lcm, c)
	}

	fmt.Println(lcm)

}

/* This is the non-LCM method.
for {
	counts++
	j := i % len(instructions)
	v := 0
	if instructions[j] == 'R' {
		v = 1
	}
	for i := range locs {
		locs[i] = network[locs[i]][v]
	}
	i++
}
*/
