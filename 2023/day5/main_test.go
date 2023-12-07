package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func getTestInput() ([]string, map[string][]SourceDestinationMap) {
	input := []string{
		"seeds: 79 14 55 13",
		// "seeds: 82 1",
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
	return seeds, maps
}

func TestApplyMap(t *testing.T) {

	seeds, maps := getTestInput()

	rs := make([]Range, 0)
	for i := 0; i < len(seeds); i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])
		rs = append(rs, Range{
			Start:  start,
			Length: length,
		})
	}

	for _, s := range []string{
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	} {
		ms := maps[s]
		rs = applyMaps(rs, ms)
		fmt.Println(rs)
	}

	// Sort the ranges by Start
	minStart := -1
	for _, r := range rs {
		if r.Start < minStart || minStart == -1 {
			minStart = r.Start
		}
	}
	fmt.Println(minStart)
}
