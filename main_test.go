package main

import "testing"

var dayOneInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
var dayThreeInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestDay1Part1(t *testing.T) {
	result := day1part1(dayOneInput)
	if result != 7 {
		t.Errorf("Expected 7, got %v", result)
	}
}

func TestDay1Part2(t *testing.T) {
	result := day1part2(dayOneInput)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestDay2Part1(t *testing.T) {
	input := make([][]string, 6)
	input[0] = []string{"forward", "5"}
	input[1] = []string{"down", "5"}
	input[2] = []string{"forward", "8"}
	input[3] = []string{"up", "3"}
	input[4] = []string{"down", "8"}
	input[5] = []string{"forward", "2"}
	result := day2part1(input)
	if result != 150 {
		t.Errorf("Expected 150, got %v", result)
	}
}

func TestDay2Part2(t *testing.T) {
	input := make([][]string, 6)
	input[0] = []string{"forward", "5"}
	input[1] = []string{"down", "5"}
	input[2] = []string{"forward", "8"}
	input[3] = []string{"up", "3"}
	input[4] = []string{"down", "8"}
	input[5] = []string{"forward", "2"}
	result := day2part2(input)
	if result != 900 {
		t.Errorf("Expected 900, got %v", result)
	}
}

func TestDay3Part1(t *testing.T) {
	result := day3part1(dayThreeInput)
	if result != 198 {
		t.Errorf("Expected 198, got %v", result)
	}
}

func TestDay3Part2(t *testing.T) {
	result := day3part2(dayThreeInput)
	if result != 230 {
		t.Errorf("Expected 230, got %v", result)
	}
}
