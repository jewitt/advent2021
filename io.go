package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func loadFields(filename string) ([][]string, error) {
	lines, err := loadLines(filename)
	if err != nil {
		return nil, err
	}

	rows := make([][]string, 0, len(lines))
	for _, line := range lines {
		rows = append(rows, strings.Fields(line))
	}

	return rows, nil
}

func loadNumbers(filename string) ([]int, error) {
	lines, err := loadLines(filename)
	if err != nil {
		return nil, err
	}

	numbers := make([]int, 0, len(lines))
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			continue // ignore non-int lines
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func loadNumbersByRegex(filename string, pattern string) ([][]int, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllSubmatch(buffer, -1)
	result := make([][]int, len(matches))
	for i, captures := range matches {
		result[i] = make([]int, len(captures) - 1)
		for j, capture := range captures[1:] {
			k, err := strconv.Atoi(string(capture))
			if err != nil {
				return nil, err
			}

			result[i][j] = k
		}
	}

	return result, nil
}

func loadLines(filename string) ([]string, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buffer), "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}
