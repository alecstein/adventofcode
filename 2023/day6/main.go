/*
Verdict: easy
Fun level: 2/10.
No major difficulty except for off-by-one errors
*/

package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getWaysToWin(d, t int) int {
	counts := 0
	for n := 0; n <= t; n++ {
		if d < n*(t-n) {
			counts++
		}
	}
	return counts
}

func getWaysToWinFast(d, t int) int {
	lowRoot := .5 * (float64(t) - math.Sqrt(float64(t*t-4*d)))
	highRoot := .5 * (float64(t) + math.Sqrt(float64(t*t-4*d)))
	// We need to count the number of integer points
	// between the roots, exclusive of the roots
	if math.Ceil(lowRoot) == lowRoot {
		lowRoot++
	} else {
		lowRoot = math.Ceil(lowRoot)
	}

	if math.Floor(highRoot) == highRoot {
		highRoot--
	} else {
		highRoot = math.Floor(highRoot)
	}

	return int(highRoot - lowRoot + 1)

}

var testInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func parseInput(lines []string) map[string][]int {
	parsedInput := make(map[string][]int, 0)
	strTimeFields := strings.Fields(lines[0])[1:]
	for _, strTime := range strTimeFields {
		intTimeField, _ := strconv.Atoi(strTime)
		parsedInput["time"] = append(parsedInput["time"], intTimeField)
	}
	strDistanceFields := strings.Fields(lines[1])[1:]
	for _, strDistance := range strDistanceFields {
		intDistanceField, _ := strconv.Atoi(strDistance)
		parsedInput["distance"] = append(parsedInput["distance"], intDistanceField)
	}
	return parsedInput
}

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/6/input")

	parsedInput := parseInput(input)

	// Part one
	waysToWin := make([]int, 0)
	for i := 0; i < len(parsedInput["time"]); i++ {
		d := parsedInput["distance"][i]
		t := parsedInput["time"][i]
		waysToWin = append(waysToWin, getWaysToWin(d, t))
	}

	total := 1
	for _, way := range waysToWin {
		total *= way
	}

	fmt.Println("The product of the number of ways to win is:")
	fmt.Println(total)

	// Part two
	// Copied this by hand
	partTwoInput := []string{
		"Time:      61677571",
		"Distance:  430103613071150",
	}
	t, _ := strconv.Atoi(strings.Fields(partTwoInput[0])[1])
	d, _ := strconv.Atoi(strings.Fields(partTwoInput[1])[1])
	fmt.Println("The number of ways to win is:")
	fmt.Println(getWaysToWinFast(d, t))
}
