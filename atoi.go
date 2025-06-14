package main

// 8. String to Integer (atoi)

func myAtoi(s string) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	sign := 1
	result := 0
	i := 0

	// Skip leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Check for overflow before multiplying and adding
		if result > (INT_MAX-digit)/10 || result < (INT_MIN+digit)/10 {
			return 0
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func myAtoi(s string) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	sign := 1
	result := 0
	i := 0

	// Skip leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Check for overflow
		if result > (INT_MAX-digit)/10 {
			if sign == 1 {
				return INT_MAX
			} else {
				return INT_MIN
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
