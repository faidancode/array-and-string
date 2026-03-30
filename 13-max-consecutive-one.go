// Given a binary array nums, return the maximum number of consecutive 1's in the array.

// Input: nums = [1,1,0,1,1,1]
// Output: 3

package arrayandstring

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0

	for _, num := range nums {
		if num == 1 {
			// Increment the current streak of 1s
			currentCount++
		} else {
			// If we hit a 0, check if the current streak was the longest
			if currentCount > maxCount {
				maxCount = currentCount
			}
			// Reset streak for the next sequence
			currentCount = 0
		}
	}

	// Final check: if the array ended with a streak of 1s,
	// we compare it one last time.
	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}

// Logic
// 1. Iterate through the array once.
// 2. If the current element is 1, increment the currentCount.
// 3. If the current element is 0, it means the streak is broken. Update maxCount if currentCount is larger, then reset currentCount to 0.
// 4. After the loop, perform one final check to ensure the last streak (if the array ends with 1s) is accounted for.
