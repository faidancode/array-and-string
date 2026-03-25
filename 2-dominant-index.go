// You are given an integer array nums where the largest integer is unique.

// Determine whether the largest element in the array is at least twice as much as every other number in the array. If it is, return the index of the largest element, or return -1 otherwise.

// Example 1:

// Input: nums = [3,6,1,0]
// Output: 1
// Explanation: 6 is the largest integer.
// For every other number in the array x, 6 is at least twice as big as x.
// The index of value 6 is 1, so we return 1.

package arrayandstring

/**
 * dominantIndex finds the index of the largest element if it is
 * at least twice as large as every other number.
 * Time Complexity: O(n) - Single pass through the array.
 * Space Complexity: O(1) - Constant space used.
 */
func dominantIndex(nums []int) int {
	maxVal := -1
	secondMax := -1
	maxIndex := -1

	for i, v := range nums {
		if v > maxVal {
			// Current max becomes the second largest before updating
			secondMax = maxVal
			maxVal = v
			maxIndex = i
		} else if v > secondMax {
			// Update second largest if current value is between max and secondMax
			secondMax = v
		}
	}

	// According to the rule: maxVal >= 2 * every other number.
	// We only need to check against the second largest to satisfy this.
	if maxVal >= 2*secondMax {
		return maxIndex
	}

	return -1
}

// Logic Breakdown

// * The "Second Max" Shortcut: You don't actually need to compare the largest number against every other number individually. If the largest number is at least twice as big as the second-largest number, it is mathematically guaranteed to be at least twice as big as all the others.
// * Initial Values: We set maxVal and secondMax to -1. Since the input numbers are between 0 and 100, any number in the array will immediately replace these placeholders.
// * Single Pass: By updating secondMax whenever a new maxVal is found, we ensure we always have the two highest values by the time the loop finishes.
