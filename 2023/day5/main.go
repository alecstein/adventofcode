/*
Oof. This was rough, and it was the first one where Part Two took me more than a day.
Even Part One took a while, just to understand what they were asking. The solution that works
for Part One won't even remotely work for Part Two: a brute-force search won't do.

The trick is to realize:
	1. Each map is basically a non-overlapping array of submaps
	2. Each submap operates piecewise on the original set of ranges to produce a new set of ranges

Visually, each application of a map does something like:
unmapped_range(s):              |-------------range--------|  |------range-------|
list_of_maps:		    { submap_A }           { submap_B }

mapped_range(s):	|--A--|         |--range--|               |------range-------|  |----B----|
Once you realize this, it just comes down to coding the solution up.
*/

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

type Range struct {
	Start  int
	Length int
}

func parseInput(input []string) ([]string, map[string][]SourceDestinationMap) {
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

func applyMaps(rs []Range, ms []SourceDestinationMap) []Range {
	mappedRs := make([]Range, 0)
	for _, m := range ms {
		unmappedRs := make([]Range, 0)
		for _, r := range rs {
			// If there's no overlap, the range is unchanged and gets added to the array
			if r.Start >= m.SourceRangeStart+m.RangeLength || r.Start+r.Length <= m.SourceRangeStart {
				unmappedRs = append(unmappedRs, r)
				continue
			}
			if (r.Start <= m.SourceRangeStart) && (r.Start+r.Length > m.SourceRangeStart) {
				// |--------range---------|    or    |--------range---|
				//        |---map----|                      |---map------|
				// |------|----------|----|          |------|------------|
				// a      c          d    b                           |xx|
				// a      c          d    b          a      c         b  d

				rLeft := r
				rLeft.Length = m.SourceRangeStart - r.Start // c - a

				rRight := r
				rRight.Start = m.SourceRangeStart + m.RangeLength                         // b
				rRight.Length = r.Start + r.Length - (m.SourceRangeStart + m.RangeLength) // b - d (can be negative)

				rMiddle := r
				rMiddle.Start = m.DestinationRangeStart // c
				rMiddle.Length = m.RangeLength
				if rRight.Length < 0 {
					rMiddle.Length += rRight.Length
				}

				if rLeft.Length > 0 {
					unmappedRs = append(unmappedRs, rLeft)
				}
				if rMiddle.Length > 0 {
					mappedRs = append(mappedRs, rMiddle)
				}

				if rRight.Length > 0 {
					unmappedRs = append(unmappedRs, rRight)
				}
			} else if (r.Start >= m.SourceRangeStart) && (r.Start < m.SourceRangeStart+m.RangeLength) {
				//          |---range---|      or           |---range------------|
				// |----------map-----------|      |----------map-----------|
				//          |-----------|xxx|               |---------------|----|

				rRight := r
				rRight.Start = m.SourceRangeStart + m.RangeLength
				rRight.Length = r.Start + r.Length - rRight.Start
				if rRight.Length < 0 {
					rRight.Length = 0
				}

				rLeft := r
				rLeft.Start = m.DestinationRangeStart + (r.Start - m.SourceRangeStart)
				rLeft.Length = r.Length - rRight.Length

				if rLeft.Length > 0 {
					mappedRs = append(mappedRs, rLeft)
				}
				if rRight.Length > 0 {
					unmappedRs = append(unmappedRs, rRight)
				}
			}
		}
		rs = unmappedRs
	}
	result := append(rs, mappedRs...)
	return result
}

func main() {

	input, _ := utils.GetPuzzleInput("https://adventofcode.com/2023/day/5/input")

	seeds, maps := parseInput(input)

	// Part one
	sourceList := make([]int, 0)
	for _, seed := range seeds {
		initSourceValue, _ := strconv.Atoi(seed)
		sourceList = append(sourceList, initSourceValue)
	}

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

	// Part two
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
	}

	// Sort the ranges by Start
	minStart := -1
	for _, r := range rs {
		if r.Start < minStart || minStart == -1 {
			minStart = r.Start
		}
	}
	fmt.Println("The minimum mapped value for part two is:")
	fmt.Println(minStart)
}
