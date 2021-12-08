package main

import (
	"math"
	"sort"
)

func day8() (int, int) {
	lines, err := loadFields("data/08/input.txt")
	if err != nil {
		panic(err)
	}

	part1count := 0
	part2count := 0

	for _, line := range lines {
		patterns := sortWords(line[:10])
		outputs := sortWords(line[11:])
		found := make([]string, 10)
		lookup := make(map[string]int, 10)
		// find 2, 3, 4 and 7
		for _, pattern := range patterns {
			switch len(pattern) {
			case 2:
				found[1] = pattern
				lookup[pattern] = 1
			case 3:
				found[7] = pattern
				lookup[pattern] = 7
			case 4:
				found[4] = pattern
				lookup[pattern] = 4
			case 7:
				found[8] = pattern
				lookup[pattern] = 8
			}
		}
		// find 6, 0, 9 and 3 - requires 4 and 1 are already found
		for _, pattern := range patterns {
			switch len(pattern) {
			case 6:
				if contains(pattern, found[1]) {
					if contains(pattern, found[4]) {
						found[9] = pattern
						lookup[pattern] = 9
					} else {
						found[0] = pattern
						lookup[pattern] = 0
					}
				} else {
					found[6] = pattern
					lookup[pattern] = 6
				}
			case 5:
				if contains(pattern, found[1]) {
					found[3] = pattern
					lookup[pattern] = 3
				}
			}
		}
		// find 5 and 2 - requires 6 is already found
		for _, pattern := range patterns {
			switch len(pattern) {
			case 5:
				if contains(found[6], pattern) {
					found[5] = pattern
					lookup[pattern] = 5
				} else if lookup[pattern] == 0 {
					found[2] = pattern
					lookup[pattern] = 2
				}
			}
		}

		for i, output := range outputs {
			part2count += lookup[output] * int(math.Pow(float64(10), float64(3 - i)))
			switch len(output) {
			case 2, 3, 4, 7:
				part1count++
			}
		}
	}

	return part1count, part2count
}

func sortWords(words []string) []string {
	for i := range words {
		words[i] = sortWord(words[i])
	}

	return words
}

func sortWord(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	return string(s)
}

func contains(haystack string, needles string) bool {
	found := 0
	for _, a := range haystack {
		for _, b := range needles {
			if a == b {
				found ++
			}
		}
	}

	return found == len(needles)
}
