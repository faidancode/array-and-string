package arrayandstring

/**
 * plusOne increments a large integer represented as a digit array.
 * Time Complexity: O(n) - Single pass through the array.
 * Space Complexity: O(n) - Only in the case where all digits are 9.
 */

func plusOne(digits []int) []int {
	n := len(digits)

	// Step 1: Traverse the array from right to left
	for i := n - 1; i >= 0; i-- {
		// If the digit is 9, it becomes 0 (carry the 1 to the next iteration)
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			// If the digit is less than 9, just add 1 and we're done
			digits[i]++
			return digits
		}

	}

	// Step 2: Handle the edge case where all digits were 9 (e.g., [9, 9] -> [1, 0, 0])
	// We create a new slice with length n+1, which defaults to all 0s.
	newDigits := make([]int, n+1)
	newDigits[0] = 1 // Set the most significant digit to 1
	return newDigits
}

// Key Takeaways
// * In-Place Modification: For most cases, we modify the digits array directly, which is very memory efficient.

// * Go Slice Optimization: Using make([]int, n+1) is cleaner than trying to "append" to the front of an existing slice, as Go's make initializes all elements to their zero-value (0) automatically.

// * Constraints: Since the array can be up to 100 digits long, we cannot convert the array to a standard integer type (like int64), because a 100-digit number is far too large for any 64-bit primitive. This manual digit manipulation is the only way to handle such large numbers.
