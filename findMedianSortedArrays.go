package main

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	numsAll := make([]int, nums1Len+nums2Len)
	copy(numsAll, nums1)
	copy(numsAll[nums1Len:], nums2)

	// Sort the merged array numsAll
	for i := range numsAll {
		for j := i + 1; j < len(numsAll); j++ {
			if numsAll[i] > numsAll[j] {
				numsAll[i], numsAll[j] = numsAll[j], numsAll[i]
			}
		}
	}

	// Calculate the median
	if (nums1Len+nums2Len)%2 == 0 {
		mid := (nums1Len + nums2Len) / 2
		return float64(numsAll[mid-1]+numsAll[mid]) / 2.0
	} else {
		mid := (nums1Len + nums2Len) / 2
		return float64(numsAll[mid])
	}
}

// func main() {
// 	nums1 := []int{1, 3}
// 	nums2 := []int{2}
// 	result := findMedianSortedArrays(nums1, nums2)
// 	println(result) // Output: 2.00000

// 	// case 2
// 	// nums1 = [1,2], nums2 = [3,4]
// 	// Output: 2.50000
// 	nums1 = []int{1, 2}
// 	nums2 = []int{3, 4}
// 	result = findMedianSortedArrays(nums1, nums2)
// 	println(result) // Output: 2.50000
// }

// findMedianSortedArrays with O(log(m+n)) time complexity
func findMedianSortedArraysOptimized(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	// Ensure nums1 is the smaller array
	if nums1Len > nums2Len {
		nums1, nums2 = nums2, nums1
		nums1Len, nums2Len = nums2Len, nums1Len
	}

	low, high := 0, nums1Len
	halfLen := (nums1Len + nums2Len + 1) / 2

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := halfLen - mid1

		if mid1 < nums1Len && nums2[mid2-1] > nums1[mid1] {
			low = mid1 + 1 // Increase mid1
		} else if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			high = mid1 - 1 // Decrease mid1
		} else {
			var maxLeft int
			if mid1 == 0 {
				maxLeft = nums2[mid2-1]
			} else if mid2 == 0 {
				maxLeft = nums1[mid1-1]
			} else {
				maxLeft = max(nums1[mid1-1], nums2[mid2-1])
			}

			if (nums1Len+nums2Len)%2 == 0 {
				var minRight int
				if mid1 == nums1Len {
					minRight = nums2[mid2]
				} else if mid2 == nums2Len {
					minRight = nums1[mid1]
				} else {
					minRight = min(nums2[mid2], nums1[mid1])
				}
				return float64(maxLeft+minRight) / 2.0
			}
			return float64(maxLeft)
		}
	}
	return 0.0
}
