// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

package arrayandstring

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // Handle cases where k > n

	// Helper function to reverse a portion of the slice
	reverse := func(s []int, start, end int) {
		for start < end {
			s[start], s[end] = s[end], s[start]
			start++
			end--
		}
	}

	// 1. Reverse entire array: [7,6,5,4,3,2,1]
	reverse(nums, 0, n-1)
	// 2. Reverse first k elements: [5,6,7, 4,3,2,1]
	reverse(nums, 0, k-1)
	// 3. Reverse remaining elements: [5,6,7, 1,2,3,4]
	reverse(nums, k, n-1)
}
