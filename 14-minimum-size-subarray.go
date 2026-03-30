// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

package arrayandstring

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	// Use n + 1 as it is impossible for a subarray to be longer than the array itself
	minLen := n + 1
	left := 0
	currentSum := 0

	for right := 0; right < n; right++ {
		currentSum += nums[right]

		// Shrink the window as much as possible while the sum is still >= target
		for currentSum >= target {
			windowSize := right - left + 1
			if windowSize < minLen {
				minLen = windowSize
			}

			// Optimization: If we find a subarray of length 1,
			// we can't get smaller than that, so we could return 1 immediately.
			if minLen == 1 {
				return 1
			}

			currentSum -= nums[left]
			left++
		}
	}

	// If minLen is still n + 1, it means no valid subarray was found
	if minLen > n {
		return 0
	}
	return minLen
}

// Logic
// 1. Initialize two pointers, left and right, both at the start of the array.
// 2. Maintain a currentSum of the elements within the window.
// 3. Expand: Move the right pointer to add elements to the currentSum until it is >= target$.
// 4. Shrink: While currentSum >= target, record the current window length (right - left + 1) and try to minimize it by moving the left pointer forward and subtracting nums[left] from the sum.
// If the result remains the initial "infinity" value, return 0.

// Summary of Logic
// 1. Expansion: The right pointer moves to include more numbers.
// 2. Contraction: The left pointer moves to "squeeze" the window to its smallest possible size while still maintaining a sum >= target.
// 3.Result: If the minLen variable was never changed from its initial value, we return 0.
