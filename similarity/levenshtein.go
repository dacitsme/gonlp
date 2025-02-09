package similarity

func LevenshteinDistance(x string, y string) int {
	if len(x) == 0 {
		return len(y)
	}
	if len(y) == 0 {
		return len(x)
	}

	lev := make([][]int, len(x)+1)
	for i := range lev {
		lev[i] = make([]int, len(y)+1)
	}

	for i := 0; i <= len(x); i++ {
		for j := 0; j <= len(y); j++ {
			if min2(i, j) == 0 {
				lev[i][j] = max(i, j)
			} else {
				deletion := lev[i-1][j] + 1
				insertion := lev[i][j-1] + 1
				substitution := lev[i-1][j-1]
				if x[i-1] != y[j-1] {
					substitution++
				}
				lev[i][j] = min3(deletion, insertion, substitution)
			}
		}
	}
	return lev[len(x)][len(y)]
}
func min3(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}
