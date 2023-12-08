package util

// Gcd calculate the greatest common denominator of two numbers.
func Gcd(a, b uint64) uint64 {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}

// Lcd calculates the least common denominator of two numbers.
func Lcd(a, b uint64) uint64 {
	return (a * b) / Gcd(a, b)
}

// LcdSum calculates the least common denominator of two fractions and sums them.
func LcdSum(n1, d1, n2, d2 uint64) (num, den uint64) {
	den = Lcd(d1, d2)

	n1 *= den / d1
	n2 *= den / d2

	num = n1 + n2

	return
}
