package main

/*
Given a string s, find the length of the longest substring without repeating characters.
Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left, maxLength := 0, 0

	for right := range len(s) {
		if index, ok := m[s[right]]; ok && index >= left {
			left = index + 1
		}
		m[s[right]] = right
		if maxLength < right-left+1 {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// Đoạn code trên định nghĩa một hàm lengthOfLongestSubstring nhận vào một chuỗi s và trả về độ dài của chuỗi con dài nhất không chứa ký tự lặp lại.
// Hàm sử dụng một bản đồ (map) để theo dõi vị trí của các ký tự trong chuỗi.
// Hai biến left và maxLength được khởi tạo để theo dõi vị trí bắt đầu của chuỗi con hiện tại và độ dài tối đa của chuỗi con không lặp lại.
// Vòng lặp for duyệt qua từng ký tự trong chuỗi s bằng cách sử dụng chỉ số right.
// Nếu ký tự hiện tại đã xuất hiện trước đó và vị trí của nó lớn hơn hoặc bằng left, thì left được cập nhật để bắt đầu từ vị trí sau ký tự lặp lại.
// Sau đó, vị trí của ký tự hiện tại được cập nhật trong bản đồ.
// Cuối cùng, nếu độ dài của chuỗi con hiện tại (right - left + 1) lớn hơn maxLength, thì maxLength được cập nhật.
// Cuối cùng, hàm trả về maxLength, độ dài của chuỗi con dài nhất không chứa ký tự lặp lại.
// Đoạn code này sử dụng một thuật toán trượt (sliding window) để tìm chuỗi con dài nhất không chứa ký tự lặp lại trong chuỗi đầu vào.
// Độ phức tạp thời gian của thuật toán là O(n), trong đó n là độ dài của chuỗi đầu vào, vì mỗi ký tự chỉ được duyệt qua một lần.
// Độ phức tạp không gian là O(min(n, m)), trong đó n là độ dài của chuỗi đầu vào và m là kích thước của bảng chữ cái (số lượng ký tự khác nhau trong chuỗi).
// Đoạn code này có thể được sử dụng để giải quyết bài toán tìm chuỗi con dài nhất không chứa ký tự lặp lại trong một chuỗi cho trước.
