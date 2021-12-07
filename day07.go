package main

import (
	"sort"
	"strconv"
	"strings"
)

func day7() (int, int) {
	lines, err := loadLines("data/07/input.txt")
	if err != nil {
		panic(err)
	}

	fields := strings.Split(string(lines[0]), ",")
	crabs := make([]int, 0, len(fields))

	for _, field := range fields {
		p, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}

		crabs = append(crabs, p)
	}

	sort.Ints(crabs)

	best := 0
	bestTriangular := 0

	for h := crabs[0]; h < crabs[len(crabs)-1]; h++ {
		fuel := 0
		fuelTriangular := 0
		for _, crab := range crabs {
			distance := crab - h
			if distance < 0 {
				distance *= -1
			}
			fuel += distance
			fuelTriangular += distance * (distance + 1) / 2
		}

		if best == 0 || fuel < best {
			best = fuel
		}
		if bestTriangular == 0 || fuelTriangular < bestTriangular {
			bestTriangular = fuelTriangular
		}
	}

	return best, bestTriangular
}
