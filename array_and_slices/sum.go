package sum

func Sum(numbers []int) int {
	var sum int

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func All(numbersToSum [][]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
