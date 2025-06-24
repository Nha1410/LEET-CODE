package main

import "fmt"

func main() {
	// Example usage of the functions
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))                                                                                                               // Should return 49
	fmt.Println(myAtoi("   -42"))                                                                                                                                        // Should return -42
	fmt.Println(isMatch("aab", "c*a*b"))                                                                                                                                 // Should return true
	fmt.Println(isPalindrome1(121))                                                                                                                                      // Should return true
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))                                                                                                           // Should return 2.0
	fmt.Println(findMedianSortedArraysOptimized([]int{1, 2}, []int{3, 4}))                                                                                               // Should return 2.5
	fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})) // Should return the sum as a linked list

	fmt.Println(convert("PAYPALISHIRING", 3))

	fmt.Println(myAtoi("   -42"))                           // Should return -42
	fmt.Println(myAtoi1("   -42"))                          // Should return -42
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))       // Should return 3
	fmt.Println(longestPalindrome("babad"))                 // Should return "bab" or "aba"
	fmt.Println(longestPalindromeExpandFromCenter("babad")) // Should return "bab" or "aba"
	fmt.Println(reverse(123))                               // Should return 321
	fmt.Println(reverse(-123))                              // Should return -321
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))             // Should return [0, 1]
}
