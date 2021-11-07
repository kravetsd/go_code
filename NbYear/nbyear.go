package nbyear

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	// your code
	percent = percent / 100
	fmt.Println("percents = ", percent, "p0 = ", p0)

	pTmp := float64(p0)
	fmt.Println("pTmp = ", pTmp, "p = ", p)
	var years = 0

	for {
		if pTmp >= float64(p) {
			return years
		} else {
			pTmp += float64(p0)*percent + float64(aug)
			fmt.Println("pTmp = ", pTmp, "p = ", p)
			years += 1
			p0 = int(pTmp)
		}
	}
}
