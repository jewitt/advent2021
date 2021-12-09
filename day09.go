package main

import (
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Basin map[Point]bool

func day9() (int, int) {
	rows, err := loadLines("data/09/input.txt")
	if err != nil {
		panic(err)
	}

	floor := make([][]int, len(rows))
	for y, cols := range rows {
		floor[y] = make([]int, len(cols))
		for x := range cols {
			i, err := strconv.Atoi(string(rows[y][x]))
			if err != nil {
				panic(err)
			}

			floor[y][x] = i
		}
	}

	risk := 0
	var basins []int
	for y, row := range floor {
		for x, height := range row {
			if y > 0 && floor[y-1][x] <= height {
				continue
			}
			if y < len(floor)-1 && floor[y+1][x] <= height {
				continue
			}
			if x > 0 && floor[y][x-1] <= height {
				continue
			}
			if x < len(floor[0])-1 && floor[y][x+1] <= height {
				continue
			}

			risk += 1 + height
			basin := make(Basin)
			findBasin(floor, Point{x, y}, basin)
			basins = append(basins, len(basin))
		}
	}

	sort.Ints(basins)
	len := len(basins)

	return risk, basins[len-1] * basins[len-2] * basins[len-3]
}

func findBasin(floor [][]int, p Point, basin Basin) {
	if p.y < 0 || p.y > len(floor)-1 || p.x < 0 || p.x > len(floor[0])-1 {
		return
	}

	if basin[p] || floor[p.y][p.x] == 9 {
		return
	}

	basin[p] = true

	findBasin(floor, Point{p.x - 1, p.y}, basin)
	findBasin(floor, Point{p.x + 1, p.y}, basin)
	findBasin(floor, Point{p.x, p.y - 1}, basin)
	findBasin(floor, Point{p.x, p.y + 1}, basin)
}
