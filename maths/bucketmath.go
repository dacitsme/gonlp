package maths

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Power(r int, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= r
	}
	return result
}

func DistinctObjectsDistinctBuckets(n, r int) int {
	return Power(r, n)
}

func IndistinctObjectsDistinctBuckets(n, r int) int {
	num := Factorial(n + r - 1)
	deno := Factorial(n) * Factorial(r-1)
	return num / deno
}
