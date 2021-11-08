package squarensum

func SquareSum(numbers []int) int {
	acc := 0
	for _, val := range numbers {
		acc += val * val
	}
	return acc
}
