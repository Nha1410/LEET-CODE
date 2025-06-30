package main

func intToRoman(num int) string {
	// Define the mapping of integers to Roman numerals
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += romans[i]
		}
	}
	return result
}

// This function use a greedy approach to convert an integer to a Roman numeral.
