package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 333AK would be
// Classes: [4, 1, 1]
// Values: [3, 14, 13]
type Hand struct {
	Cards   string
	Bid     int
	Classes []int
	Values  []int
}

var cardOrder = map[rune]int{
	'A': 1,
	'K': 2,
	'Q': 3,
	'J': 4,
	'T': 5,
	'9': 6,
	'8': 7,
	'7': 8,
	'6': 9,
	'5': 10,
	'4': 11,
	'3': 12,
	'2': 13,
}

// Output a list of card hands and their bids
func parseInput(input []string) (hands []Hand) {
	for _, line := range input {
		if line != "" {
			fields := strings.Fields(line)
			cards := fields[0]
			bid, _ := strconv.Atoi(fields[1])
			hands = append(hands, Hand{
				Cards: cards,
				Bid:   bid,
			})
		}
	}
	return hands
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func classifyHand(h *Hand) {
	sortedHand := sortString(h.Cards)
	currentChar := sortedHand[0]
	currentCount := 1
	for _, r := range sortedHand[1:] {
		if r == rune(currentChar) {
			currentCount++
		} else {
			h.Classes = append(h.Classes, currentCount)
			h.Values = append(h.Values, cardOrder[r])
			currentChar = byte(r)
			currentCount = 1
		}
	}

}

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/7/input")
	hands := parseInput(input)

	classifyHand(&hands[0])

	// compareHands := func(i, j int) bool {
	// 	handA := cardStringToRepresentation(handsAndBids[i][0])
	// 	handB := cardStringToRepresentation(handsAndBids[j][0])
	// 	_, handAClassRank := getCardClass(handA)
	// 	_, handBClassRank := getCardClass(handB)

	// 	return handAClassRank > handBClassRank // Adjust this comparison as needed
	// }

	// sort.Slice(handsAndBids, compareHands)

	fmt.Println(hands[0])
}
