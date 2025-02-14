package maths

func FactorialFunction(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialFunction(n-1)
}

func Permutation(s string) int {
	n := len(s)
	Countchar := make(map[rune]int)
	for _, char := range s {
		Countchar[char]++
	}

	numerator := FactorialFunction(n)
	denominator := 1
	for _, count := range Countchar {
		denominator *= FactorialFunction(count)
	}

	return numerator / denominator

}
