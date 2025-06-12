package main

func longestPalindrome(s string) string {
	maxLength := 0
	res := ""

	for i := range len(s) {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && (j-i) > maxLength {
				maxLength = j - i
				res = s[i:j]
			}
		}
	}

	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// tôi muốn triển khai theo cách mở rộng từ trung tâm đến hai đầu
func longestPalindromeExpandFromCenter(s string) string {
	maxLength := 0
	res := ""

	for i := 0; i < len(s); i++ {
		// Kiểm tra chuỗi đối xứng với tâm là ký tự đơn
		len1 := expandFromCenter(s, i, i)
		if len1 > maxLength {
			maxLength = len1
			res = s[i-(len1-1)/2 : i+len1/2+1]
		}

		// Kiểm tra chuỗi đối xứng với tâm là khoảng giữa hai ký tự
		len2 := expandFromCenter(s, i, i+1)
		if len2 > maxLength {
			maxLength = len2
			res = s[i-(len2-1)/2 : i+len2/2+1]
		}
	}

	return res
}
func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1 // Trả về độ dài của chuỗi đối xứng
}
