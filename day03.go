package main

import (
	"strconv"
)

func day3() (int, int) {
	lines, err := loadLines("data/03/input.txt")
	if err != nil {
		panic(err)
	}

	return day3part1(lines), day3part2(lines)
}

func day3part1(lines []string) int {
	bits := make([]int, len(lines[0]))
	for _, line := range lines {
		for i := range line {
			value, err := strconv.Atoi(string(line[i]))
			if err == nil {
				bits[i] += value
			}
		}
	}

	var gamma, epsilon string
	for _, bit := range bits {
		if bit >= len(lines)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(g * e)
}

func day3part2(lines []string) int {
	co2 := make([]string, len(lines))
	oxygen := make([]string, len(lines))
	copy(co2, lines)
	copy(oxygen, lines)

	return day3part2process(co2, true) * day3part2process(oxygen, false)
}

func day3part2process(lines []string, isCo2 bool) int {
	p := 0
	for len(lines) != 1 {
		zeroes := 0
		ones := 0
		for _, line := range lines {
			if line[p] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		var bit byte = '0'
		if isCo2 && zeroes > ones || !isCo2 && zeroes <= ones {
			bit = '1'
		}

		for i := 0; i < len(lines); i++ {
			if lines[i][p] == bit {
				lines = append(lines[:i], lines[i+1:]...)
				i--
			}
		}

		p++
	}

	v, _ := strconv.ParseInt(lines[0], 2, 54)

	return int(v)
}
