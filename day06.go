package main

import(
	"strconv"
	"strings"
)

func day6() (int, int) {
	lines, err := loadLines("data/06/input.txt")
	if err != nil {
		panic(err)
	}

	return countFishes(lines, 80), countFishes(lines, 256)
}

func countFishes(lines []string, days int) (total int) {
	fishes := make([]int, 9)
	for _, line := range lines {
		fields := strings.Split(string(line), ",")
		for _, field := range fields {
			state, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}

			fishes[state]++
		}
	}

	for day := 0; day < days; day++ {
		total = 0
		newfish := 0
		for state, count := range fishes {
			if state == 0 {
				newfish = count
			} else {
				fishes[state - 1] = count
				total += count
			}
		}
		fishes[6] += newfish
		fishes[8] = newfish
		total += newfish*2
	}

	return total
}
