package grow

func Grow(arr []int) int {
	acc := 1
	for _, val := range arr {
		acc *= val
	}
	return acc
}
