package main

/*
I found this problem pretty tedious, but I do feel like it exposed some
of my weaknesses in string/struct manipulation.
*/

import (
	"adventofcode/utils"
	"fmt"
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

// Takes in a string of game data, processes it, and puts it into a Game struct
func marshalDataIntoGame(data string) Game {
	gameRecordSplit := strings.Split(data, ": ")

	gameId, _ := strconv.Atoi(strings.Split(gameRecordSplit[0], " ")[1])
	records := make([]Record, 0)
	g := Game{
		ID:      gameId,
		Records: records,
	}

	recordStrings := strings.Split(gameRecordSplit[1], "; ")
	for _, recordString := range recordStrings {
		intColorStrings := strings.Split(recordString, ", ")
		record := Record{}
		for _, intColorString := range intColorStrings {
			intColorArray := strings.Split(intColorString, " ")
			switch intColorArray[1] {
			case "red":
				record.Red, _ = strconv.Atoi(intColorArray[0])
			case "green":
				record.Green, _ = strconv.Atoi(intColorArray[0])
			case "blue":
				record.Blue, _ = strconv.Atoi(intColorArray[0])
			}
		}
		g.Records = append(g.Records, record)
	}
	return g
}

func (g Game) PossibleGames(redLimit, greenLimit, blueLimit int) bool {
	for _, record := range g.Records {
		if record.Red > redLimit || record.Green > greenLimit || record.Blue > blueLimit {
			return false
		}
	}
	return true
}

func (g Game) MinSet() (int, int, int) {
	var minRed int
	var minBlue int
	var minGreen int

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
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/2/input")

	gameStrings := strings.Split(input, "\n")
	games := make([]Game, 0)
	for _, gameString := range gameStrings {
		if gameString == "" {
			continue
		}
		games = append(games, marshalDataIntoGame(gameString))
	}

	var s int
	for _, g := range games {

		if g.PossibleGames(12, 13, 14) {
			s += g.ID
		}
	}
	fmt.Printf("Sum of IDs of possible games using %d red, %d green, and %d blue cubes: %d\n", 12, 13, 14, s)

	var p int
	for _, g := range games {
		minRed, minBlue, minGreen := g.MinSet()
		cubePower := minRed * minBlue * minGreen
		p += cubePower
	}
	fmt.Printf("Sum of cube powers: %d\n", p)

}
