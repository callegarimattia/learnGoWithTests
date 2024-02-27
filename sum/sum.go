package sum

func Sum(numbers []int) int {
	var sum int

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func All(numbersToSum [][]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func AllUnoptimized(numbersToSum [][]int) []int {
	var sums []int //nolint: prealloc

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}
