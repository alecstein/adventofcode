package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type SourceDestinationMap struct {
	SourceRangeStart      int
	DestinationRangeStart int
	RangeLength           int
}

func main() {

	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/5/input")

	input = []string{
		// "seeds: 14",
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	// Parse the puzzle input
	var seeds []string
	maps := make(map[string][]SourceDestinationMap)

	split := func(r rune) bool { return r == ':' }
	var k string
	for _, line := range input {
		if strings.Contains(line, ":") {
			unparsedK := strings.FieldsFunc(line, split)
			k = unparsedK[0]
			if len(unparsedK) == 2 {
				// The first line needs to be handled differently
				seeds = strings.Fields(unparsedK[1])
			}
		} else if line != "" {
			fields := strings.Fields(line)
			destinationRangeStart, _ := strconv.Atoi(fields[0])
			sourceRangeStart, _ := strconv.Atoi(fields[1])
			rangeLength, _ := strconv.Atoi(fields[2])
			maps[k] = append(maps[k], SourceDestinationMap{
				sourceRangeStart,
				destinationRangeStart,
				rangeLength},
			)
		}
	}

	// Part one
	sourceList := make([]int, 0)
	for _, seed := range seeds {
		initSourceValue, _ := strconv.Atoi(seed)
		sourceList = append(sourceList, initSourceValue)
	}
	fmt.Println(sourceList)
	for _, mapName := range []string{
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	} {
		destinationList := make([]int, 0)
		for _, sourceValue := range sourceList {
			mapped := false
			for _, m := range maps[mapName] {
				if (m.SourceRangeStart <= sourceValue) && (sourceValue < m.SourceRangeStart+m.RangeLength) {
					destinationValue := sourceValue - (m.SourceRangeStart - m.DestinationRangeStart)
					destinationList = append(destinationList, destinationValue)
					mapped = true
					break
				}
			}
			if !mapped {
				destinationList = append(destinationList, sourceValue)
			}
		}
		sourceList = destinationList
		fmt.Println(sourceList)
	}

	minValue := -1
	for _, v := range sourceList {
		if v < minValue || minValue == -1 {
			minValue = v
		}
	}

	// fmt.Println(sourceDestinationMap)
	fmt.Println("The minimum mapped value for part one is:")
	fmt.Println(minValue)

	// 	for _, initSourceValue := range sourceList {
	//  {
	// 			for _, row := range maps[mapName] {
	// 				destinationRangeStart, _ := strconv.Atoi(row[0])
	// 				sourceRangeStart, _ := strconv.Atoi(row[1])
	// 				rangeLength, _ := strconv.Atoi(row[2])
	// 				currentSourceVaue := sourceList[initSourceValue]
	// 				if (sourceRangeStart <= currentSourceVaue) && (currentSourceVaue < sourceRangeStart+rangeLength) {
	// 					destinationValue := currentSourceVaue - (sourceRangeStart - destinationRangeStart)
	// 					sourceList[initSourceValue] = destinationValue
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}

	// minMappedValue := -1
	// for _, v := range sourceList {
	// 	if v < minMappedValue || minMappedValue == -1 {
	// 		minMappedValue = v
	// 	}
	// }

	// fmt.Println(sourceDestinationMap)
	// fmt.Println("The minimum mapped value for part one is:")
	// fmt.Println(sourceList)
	// fmt.Println(minMappedValue)

	// sourceList := make([]int, 0)
	// var initSourceValue int
	// for i, seed := range seeds {
	// 	if i%2 == 0 {
	// 		initSourceValue, _ = strconv.Atoi(seed)
	// 	} else {
	// 		rangeValue, _ := strconv.Atoi(seed)
	// 		for i := initSourceValue; i < initSourceValue+rangeValue; i++ {
	// 			sourceList = append(sourceList, i)
	// 		}
	// 	}
	// }

	// fmt.Println(len(sourceList))

	// fmt.Println(sourceDestinationMap)

	// for initSourceValue, _ := range sourceDestinationMap {
	// 	for _, mapName := range []string{
	// 		"seed-to-soil map",
	// 		"soil-to-fertilizer map",
	// 		"fertilizer-to-water map",
	// 		"water-to-light map",
	// 		"light-to-temperature map",
	// 		"temperature-to-humidity map",
	// 		"humidity-to-location map",
	// 	} {
	// 		for _, row := range maps[mapName] {
	// 			destinationRangeStart, _ := strconv.Atoi(row[0])
	// 			sourceRangeStart, _ := strconv.Atoi(row[1])
	// 			rangeLength, _ := strconv.Atoi(row[2])
	// 			currentSourceVaue := sourceDestinationMap[initSourceValue]
	// 			if (sourceRangeStart <= currentSourceVaue) && (currentSourceVaue < sourceRangeStart+rangeLength) {
	// 				destinationValue := currentSourceVaue - (sourceRangeStart - destinationRangeStart)
	// 				sourceDestinationMap[initSourceValue] = destinationValue
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// minMappedValue = -1
	// for _, v := range sourceDestinationMap {
	// 	if v < minMappedValue || minMappedValue == -1 {
	// 		minMappedValue = v
	// 	}
	// }

	// fmt.Println("The minimum mapped value for part two is:")
	// fmt.Println(minMappedValue)
}
