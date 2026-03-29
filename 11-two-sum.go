// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

// Return the indices of the two numbers index1 and index2, each incremented by one, as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

package arrayandstring

// TwoSum finds two numbers that add up to a specific target.
// It returns the 1-indexed positions of these numbers.
func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			// Found the solution; return 1-indexed results
			return []int{left + 1, right + 1}
		}

		if currentSum < target {
			// Sum is too low, move left pointer to a larger value
			left++
		} else {
			// Sum is too high, move right pointer to a smaller value
			right--
		}
	}

	// The problem guarantees exactly one solution,
	// so we don't need to handle the "not found" case explicitly.
	return []int{}
}
