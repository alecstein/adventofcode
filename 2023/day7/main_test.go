package main

// import (
// 	"fmt"
// 	"sort"
// 	"testing"
// )

// func TestCardStringToRepresentation(t *testing.T) {
// 	fmt.Println(cardStringToRepresentation("QKKA2"))
// }

// func TestGetCardClass(t *testing.T) {
// 	testCases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{"KKKKK", "five of a kind"},
// 		{"KKAKK", "four of a kind"},
// 		{"KKQKQ", "full house"},
// 		{"TTT98", "three of a kind"},
// 		{"TT998", "two pair"},
// 		{"AA123", "one pair"},
// 		{"AKQJT", "high card"},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(fmt.Sprintf("getCardClass(%s)", tc.input), func(t *testing.T) {
// 			actual, _ := getCardClass(cardStringToRepresentation(tc.input))
// 			if actual != tc.expected {
// 				t.Errorf("getCardClass(%s) expected %s, got %s", tc.input, tc.expected, actual)
// 			}
// 		})
// 	}
// }

// var customOrder = map[rune]int{
// 	'A': 1,
// 	'K': 2,
// 	'Q': 3,
// 	'J': 4,
// 	'T': 5,
// 	'9': 6,
// 	'8': 7,
// 	'7': 8,
// 	'6': 9,
// 	'5': 10,
// 	'4': 11,
// 	'3': 12,
// 	'2': 13,
// }

// func orderStringByRank(s string) string {
// 	r := []rune(s)
// 	sort.Slice(r,
// 		func(i, j int) bool {
// 			return customOrder[r[i]] < customOrder[r[j]]
// 		})
// 	return string(r)
// }

// func TestOrderStringByRank(t *testing.T) {
// 	testCases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{"KKKKK", "KKKKK"},
// 		{"22987", "98722"},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(fmt.Sprintf("orderStringByRank(%s)", tc.input), func(t *testing.T) {
// 			actual := orderStringByRank(tc.input)
// 			if actual != tc.expected {
// 				t.Errorf("orderStringByRank(%s) expected %s, got %s", tc.input, tc.expected, actual)
// 			}
// 		})
// 	}
// }

// func cardStringToRepresentation2(s string) map[string]int {
// 	r := make(map[string]int)
// 	for _, card := range s {
// 		r[string(card)] += 1
// 	}

// 	// Now we want to create a new representation
// 	// based on this dictionary: we sort by value
// 	// then by key and then join them back as a string

// }
