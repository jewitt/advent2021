package main

func day1() (int, int) {
	numbers, err := loadNumbers("data/01/input.txt")
	if err != nil {
		panic(err)
	}

	return day1part1(numbers), day1part2(numbers)
}

func day1part1(numbers []int) (count int) {
	for i, number := range numbers {
		if i > 0 && number > numbers[i-1] {
			count++
		}
	}

	return count
}

func day1part2(numbers []int) (count int) {
	for i := range numbers {
		if i > 1 && i < len(numbers)-1 {
			if numbers[i-2]+numbers[i-1]+numbers[i] < numbers[i-1]+numbers[i]+numbers[i+1] {
				count++
			}
		}
	}

	return count
}
