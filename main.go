package main

import (
	"fmt"
)

func main() {
	days := []func() (int, int){
		day1,
		day2,
		day3,
		day4,
		day5,
		day6,
		day7,
	}

	fmt.Printf("running...\n\n")
	for day, fn := range days {
		p1, p2 := fn()
		fmt.Printf("  day %d part 1: %d\n  day %d part 2: %d\n\n", day+1, p1, day+1, p2)
	}
	fmt.Println("done")
}
