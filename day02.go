package main

import (
	"fmt"
	"strconv"
)

func day2() (int, int) {
	fields, err := loadFields("data/02/input.txt")
	if err != nil {
		panic(err)
	}

	return day2part1(fields), day2part2(fields)
}

func day2part1(rows [][]string) int {
	var horizontal, vertical int
	for _, row := range rows {
		distance, err := strconv.Atoi(row[1])
		if err != nil {
			panic(fmt.Sprintf("invalid number in %v\n", row))
		}

		switch row[0] {
		case "forward":
			horizontal += distance
		case "up":
			vertical -= distance
		case "down":
			vertical += distance
		}
	}

	return horizontal * vertical
}

func day2part2(rows [][]string) int {
	var horizontal, vertical, aim int
	for _, row := range rows {
		distance, err := strconv.Atoi(row[1])
		if err != nil {
			panic(fmt.Sprintf("invalid number in %v\n", row))
		}

		switch row[0] {
		case "forward":
			horizontal += distance
			vertical += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	return horizontal * vertical
}
