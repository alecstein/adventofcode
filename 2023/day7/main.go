package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards      string
	Bid        int
	SortedHand map[int][]int
	Class      int
}

var cardOrderPartOne = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardOrderPartTwo = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 11,
	'9': 10,
	'8': 9,
	'7': 8,
	'6': 7,
	'5': 6,
	'4': 5,
	'3': 4,
	'2': 3,
	'J': 2,
}

// Output a list of card hands and their bids
func parseInput(input []string) (hands []Hand) {
	for _, line := range input {
		if line != "" {
			fields := strings.Fields(line)
			bid, _ := strconv.Atoi(fields[1])
			hands = append(hands, Hand{
				Cards: fields[0],
				Bid:   bid,
			})
		}
	}
	return hands
}

// func sortString(s string) string {
// 	r := []rune(s)
// 	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
// 	return string(r)
// }

func scoreHand(h *Hand) map[int][]int {

	// Maps AAQQ9
	// to {
	// 	A: 2,
	// 	Q: 2,
	// 	9: 1,
	// }
	charCounts := make(map[rune]int, 0)
	for _, r := range h.Cards {
		charCounts[r] += 1
	}

	// Maps {
	// 	2: [12, 14],
	// 	1: [9],
	// }
	sortedHand := make(map[int][]int, 0)
	for k, v := range charCounts {
		sortedHand[v] = append(sortedHand[v], cardOrderPartOne[k])
	}

	// Maps {
	// 	2: [14, 12],
	// 	1: [9],
	// }
	for _, v := range sortedHand {
		sort.Slice(v, func(i, j int) bool { return v[i] > v[j] })
	}

	return sortedHand
}

func classifyHand(sortedHand map[int][]int) int {
	if len(sortedHand[5]) != 0 {
		return 7
	} else if len(sortedHand[4]) != 0 {
		return 6
	} else if len(sortedHand[3]) != 0 && len(sortedHand[2]) != 0 {
		return 5
	} else if len(sortedHand[3]) != 0 {
		return 4
	} else if len(sortedHand[2]) == 2 {
		return 3
	} else if len(sortedHand[2]) == 1 {
		return 2
	} else {
		return 1
	}
}

func (h *Hand) classifyHandPartTwo() int {
	numJs := 0
	for _, v := range h.Cards {
		if v == 'J' {
			numJs++
		}
	}
	if len(h.SortedHand[5]) != 0 {
		// 5 of a kind
		return 7 // 5 of a kind
	} else if len(h.SortedHand[4]) != 0 {
		if numJs > 0 {
			return 7 // five of a kind
		}
		return 6 // four of a kind
	} else if len(h.SortedHand[3]) != 0 && len(h.SortedHand[2]) != 0 {
		// full house
		if numJs == 2 || numJs == 3 {
			return 7 // five of a kind
		}
		return 5 // full house
	} else if len(h.SortedHand[3]) != 0 {
		// 3 of a kind
		if numJs == 1 || numJs == 3 {
			return 6 // four of a kind
		}
		return 4 // 3 of a kind
	} else if len(h.SortedHand[2]) == 2 {
		if numJs == 2 {
			return 6 // four of a kind
		} else if numJs == 1 {
			return 5 // full house
		}
		return 3 // two pair
	} else if len(h.SortedHand[2]) == 1 {
		// one pair
		if numJs == 1 || numJs == 2 {
			return 4
		}
		return 2 // one pair
	} else {
		if numJs == 1 {
			return 2
		}
		return 1
	}
}

func main() {
	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/7/input")
	// input = []string{
	// 	"32T3K 765",
	// 	"T55J5 684",
	// 	"KK677 28",
	// 	"KTJJT 220",
	// 	"QQQJA 483",
	// }

	// input = []string{
	// 	"JJJ8J 1",
	// 	"JJJJJ 2",
	// 	"JJ222 3",
	// 	"J88J8 4",
	// 	"8JTJT 12",
	// 	"8222J 11",
	// 	"8J888 13",
	// }
	hands := parseInput(input)

	for i, h := range hands {
		hands[i].SortedHand = scoreHand(&h)
		hands[i].Class = classifyHand(hands[i].SortedHand)
	}

	sort.Slice(hands, func(i, j int) bool {
		// Sort by the first-card rule
		if hands[i].Class != hands[j].Class {
			return hands[i].Class > hands[j].Class
		} else {
			for k := 0; k < 5; k++ {
				if cardOrderPartOne[rune(hands[i].Cards[k])] != cardOrderPartOne[rune(hands[j].Cards[k])] {
					return cardOrderPartOne[rune(hands[i].Cards[k])] > cardOrderPartOne[rune(hands[j].Cards[k])]
				}
			}
		}
		return true
	})

	total := 0
	for i, h := range hands {
		rank := len(hands) - i
		total += h.Bid * rank
	}

	fmt.Println("The answer to part one is:")
	fmt.Println(total)

	for i := range hands {
		hands[i].Class = hands[i].classifyHandPartTwo()
	}

	sort.Slice(hands, func(i, j int) bool {
		// Sort by the first-card rule
		if hands[i].Class != hands[j].Class {
			return hands[i].Class > hands[j].Class
		} else {
			for k := 0; k < 5; k++ {
				cardIValue := cardOrderPartTwo[rune(hands[i].Cards[k])]
				cardJValue := cardOrderPartTwo[rune(hands[j].Cards[k])]
				if cardIValue != cardJValue {
					return cardIValue > cardJValue
				}
			}
		}
		return true
	})

	total = 0
	for i, h := range hands {
		rank := len(hands) - i
		// for _, v := range h.Cards {
		// 	if v == 'J' {
		// 		fmt.Println(h.Cards, h.SortedHand, h.Bid, rank, h.Class)
		// 		break
		// 	}

		// }
		total += h.Bid * rank
	}

	fmt.Println("The answer to part two is: ")
	fmt.Println(total)

}

// // 5, 4, 3/2, 3, 2/2, 1
// sort.Slice(hands, func(i, j int) bool {
// 	if hands[i].Class != hands[j].Class {
// 		return hands[i].Class > hands[j].Class
// 	} else {
// 		for k := 7; k > 0; k-- {
// 			for i, v := range hands[i].SortedHand[k] {
// 				if v != hands[j].SortedHand[k][i] {
// 					return v > hands[j].SortedHand[k][i]
// 				}
// 			}

// 		}
// 	}
// 	return false
// })

// compareHands := func(i, j int) bool {
// 	handA := cardStringToRepresentation(handsAndBids[i][0])
// 	handB := cardStringToRepresentation(handsAndBids[j][0])
// 	_, handAClassRank := getCardClass(handA)
// 	_, handBClassRank := getCardClass(handB)

// 	return handAClassRank > handBClassRank // Adjust this comparison as needed
// }

// sort.Slice(handsAndBids, compareHands)

// fmt.Println(hands[0])

// classifyHand(&hands[0])
