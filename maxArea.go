package main

// 11. Container With Most Water
// ✅ Two Pointers (Hai con trỏ)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		h := min(height[left], height[right])
		area := width * h
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// my Code, bad time complexity O(n^2)
// func maxArea(height []int) int {
// 	maxS := 0

// 	for i := 0; i < len(height); i++ {
// 		for j := i; j < len(height); j++ {
// 			minHeight := min(height[i], height[j])

// 			S := minHeight * (j - i)
// 			if S > maxS {
// 				maxS = S
// 			}

// 		}
// 	}

// 	return maxS
// }
