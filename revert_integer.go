package main

// 7. Reverse Integer
// Medium
// Topics
// premium lock icon
// Companies
// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func reverse(x int) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before multiplying and adding
		if result > (INT_MAX-digit)/10 || result < (INT_MIN-digit)/10 {
			return 0
		}

		result = result*10 + digit
	}

	return result
}
