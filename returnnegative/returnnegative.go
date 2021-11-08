package returnnegative

func MakeNegative(x int) int {
	switch {
	case x > 0:
		return -x
	case x < 0:
		return x
	default:
		return 0
	}
}

func MakeNegativeBestPractive(x int) int {
	if x >= 0 {
		return -x
	}
	return x
}
