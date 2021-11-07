package reversestrings

func Solution(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Solutionbestpractice(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return result
}
