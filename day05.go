package main

func day5() (int, int) {
	matches, err := loadNumbersByRegex("data/05/input.txt", "(\\d+),(\\d+)\\s+\\->\\s+(\\d+),(\\d+)")
	if err != nil {
		panic(err)
	}

	return day5process(matches, false), day5process(matches, true) 
}

func day5process(matches [][]int, diagonals bool) int {
	grid := make([][]int, 1000)
	for y := range grid {
		grid[y] = make([]int, 1000)
	}

	for _, m := range matches {
		sx, sy, tx, ty := m[0], m[1], m[2], m[3]
		if !diagonals && sx != tx && sy != ty {
			continue
		}
		var dx, dy int
		if sx < tx {
			dx = 1
		}
		if sx > tx {
			dx = -1
		}
		if sy < ty {
			dy = 1
		}
		if sy > ty {
			dy = -1
		}

		x, y := sx, sy
		for {
			grid[y][x]++
			if x == tx && y == ty {
				break
			}
			x += dx
			y += dy
		}
	}

	count := 0
	for _, y := range grid {
		for _, x := range y {
			if x >= 2 {
				count++
			}
		}
	}

	return count
}
