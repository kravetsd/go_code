package nbyear

func NbYear(p0 int, percent float64, aug int, p int) int {
	// your code
	percent = percent / 100

	pTmp := float64(p0)
	var years = 0

	for {
		if pTmp >= float64(p) {
			return years
		} else {
			pTmp += float64(p0)*percent + float64(aug)
			years += 1
			p0 = int(pTmp)
		}
	}
}

func NbYearBestPractices(p0 int, percent float64, aug int, p int) (n int) {
	if p0 >= p {
		return n
	}
	newP := p0 + aug + int(float64(p0)*percent/100)
	n++
	return n + NbYear(newP, percent, aug, p)
}
