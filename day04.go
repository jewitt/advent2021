package main

import (
	"strconv"
	"strings"
)

type card struct {
	numbers [][]string
	marked  [][]bool
	called  bool
}

func (c *card) mark(number string) bool {
	if c.completed() {
		return false
	}

	for i, cols := range c.numbers {
		for j := range cols {
			if c.numbers[i][j] == number {
				c.marked[i][j] = true

				return true
			}
		}
	}

	return false
}

func (c *card) completed() bool {
	if c.called {
		return false // if we already called we're out of the game
	}

	for i := 0; i < 5; i++ {
		var markedRow, markedCol int
		for j := 0; j < 5; j++ {
			if c.marked[i][j] {
				markedRow++
			}
			if c.marked[j][i] {
				markedCol++
			}

			if markedCol == 5 || markedRow == 5 {
				c.called = true

				return true
			}
		}
	}

	return false
}

func (c *card) score(multiplier int) int {
	sum := 0
	for i, cols := range c.numbers {
		for j := range cols {
			if c.marked[i][j] == false {
				value, _ := strconv.Atoi(c.numbers[i][j])
				sum += value
			}
		}
	}

	return multiplier * sum
}

func newCard() *card {
	var card card
	card.numbers = make([][]string, 5)
	for k := range card.numbers {
		card.numbers[k] = make([]string, 5)
	}

	card.marked = make([][]bool, 5)
	for k := range card.marked {
		card.marked[k] = make([]bool, 5)
	}

	return &card
}

func day4() (int, int) {
	lines, err := loadLines("data/04/input.txt")
	if err != nil {
		panic(err)
	}

	var first, last int

	randoms := strings.Split(lines[0], ",")

	var cards []*card
	for i := 2; i < len(lines); i += 6 {
		card := newCard()
		for k := 0; k < 5; k++ {
			for j, number := range strings.Fields(lines[i+k]) {
				card.numbers[k][j] = number
			}
		}
		cards = append(cards, card)
	}

	for _, random := range randoms {
		for _, card := range cards {
			if card.mark(random) && card.completed() {
				multiplier, _ := strconv.Atoi(random)

				score := card.score(multiplier)

				if score != 0 {
					if first == 0 {
						first = score
					}
					last = score
				}
			}
		}
	}

	return first, last
}
