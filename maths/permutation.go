package maths

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Permutation(s string) int {
	n := len(s)
	Countchar := make(map[rune]int)
	for _, char := range s {
		Countchar[char]++
	}

	numerator := Factorial(n)
	denominator := 1
	for _, count := range Countchar {
		denominator *= Factorial(count)
	}

	return numerator / denominator

}
