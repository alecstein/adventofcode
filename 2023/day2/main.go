package main

/*
I found this problem pretty tedious, but I do feel like it exposed some
of my weaknesses in string/struct manipulation.
*/

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID      int
	Records []Record
}

// Pull out the game ID from line
func getGameIdString(line string) string {
	re := regexp.MustCompile(`Game (\d+)`)
	match := re.FindStringSubmatch(line)
	return match[1]
}

// Pull out a list of records from line. Returns
// [
//
//	1 red, 2 green, 3 blue,
//	4 red, 5 green, 6 blue,
//
// ] etc
func getRecordStrings(line string) []string {
	gameRecordSplit := strings.Split(line, ": ")
	recordStrings := strings.Split(gameRecordSplit[1], "; ")
	return recordStrings
}

// Takes in a string of game data, processes it, and puts it into a Game struct
func lineToStruct(line string) Game {
	game := Game{}
	game.ID, _ = strconv.Atoi(getGameIdString(line))

	game.Records = make([]Record, 0)

	for _, recordString := range getRecordStrings(line) {
		record := Record{}
		re := regexp.MustCompile(`(\d+) (\w+)`)
		matches := re.FindAllStringSubmatch(recordString, -1)
		for _, match := range matches {
			switch match[2] {
			case "red":
				record.Red, _ = strconv.Atoi(match[1])
			case "green":
				record.Green, _ = strconv.Atoi(match[1])
			case "blue":
				record.Blue, _ = strconv.Atoi(match[1])
			}
		}
		game.Records = append(game.Records, record)
	}
	return game
}

// Checks if a game is possible given the limits of red, green, and blue cubes
func (g Game) isPossible(redLimit, greenLimit, blueLimit int) bool {
	for _, record := range g.Records {
		if record.Red > redLimit || record.Green > greenLimit || record.Blue > blueLimit {
			return false
		}
	}
	return true
}

func (g Game) minSet() (int, int, int) {
	var minRed, minBlue, minGreen int

	for _, record := range g.Records {
		if record.Red > minRed {
			minRed = record.Red
		}
		if record.Green > minGreen {
			minGreen = record.Green
		}
		if record.Blue > minBlue {
			minBlue = record.Blue
		}

	}
	return minRed, minBlue, minGreen
}

func main() {
	lines, _ := utils.GetPuzzleInputAsLines("https://adventofcode.com/2023/day/2/input")

	// Part one
	s := 0
	games := make([]Game, 0)
	for _, line := range lines {
		games = append(games, lineToStruct(line))
	}

	for _, game := range games {
		if game.isPossible(12, 13, 14) {
			s += game.ID
		}
	}
	fmt.Printf("Sum of IDs of possible games using %d red, %d green, and %d blue cubes:\n%d\n", 12, 13, 14, s)

	// Part two
	s = 0
	for _, game := range games {
		minRed, minBlue, minGreen := game.minSet()
		cubePower := minRed * minBlue * minGreen
		s += cubePower
	}
	fmt.Printf("Sum of cube powers:\n%d\n", s)

}
