package main

func romanToInt(s string) int {
	// Define the mapping of Roman numerals to integers
	romanToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanToIntMap[s[i]]

		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}

		prevValue = currentValue
	}

	return total
}

// This function uses a greedy approach to convert a Roman numeral to an integer.
