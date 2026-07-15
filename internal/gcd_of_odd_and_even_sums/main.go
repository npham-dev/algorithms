package gcd_of_odd_and_even_sums

func gcdOfOddEvenSums(n int) int {
	return gcd(n*(n+1), n*n)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
