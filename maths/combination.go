package maths

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

func Combination(n int, k int) int {
	numerator := Fact(n)
	denominator := Fact(k) * Fact(n-k)
	return numerator / denominator

}
