package calculator

// Fib Calculates the Nth fibonacci number
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	c := 1
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

// Mdc calculates the lcm of two numbers
func Mdc(a, b int) int {
	if b == 0 {
		return a
	}
	return Mdc(b, a%b)
}

// Mmc calculates the hcm of two numbers
func Mmc(a, b int) int {
	return a * b / Mdc(a, b)
}

// Pow Calculates base^(exponent)
func Pow(base, exponent int) int {
	ans := 1
	if exponent == 0 {
		return 1
	}
	if exponent < 0 {
		base = 1 / base
		exponent *= -1
	}
	for exponent > 0 {
		if exponent%2 == 1 {
			ans *= base
		}
		base = base * base
		exponent /= 2
	}
	return ans
}
