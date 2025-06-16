package main

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	// Calculate the divisor to extract the most significant digit
	divisor := 1
	for x/divisor >= 10 {
		divisor *= 10
	}

	for x > 0 {
		left := x / divisor
		right := x % 10

		if left != right {
			return false
		}

		// Remove the most significant and least significant digits
		x = (x % divisor) / 10
		divisor /= 100
	}

	return true
}
