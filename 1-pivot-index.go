package arrayandstring

/**
 * pivotIndex calculates the leftmost pivot index.
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */
func pivotIndex(nums []int) int {
	totalSum := 0
	// Step 1: Calculate the sum of all elements in the array
	for _, val := range nums {
		totalSum += val
	}

	leftSum := 0
	// Step 2: Iterate through the array to find the pivot
	for i, val := range nums {
		// Logic:
		// rightSum = totalSum - leftSum - currentVal
		// We check if leftSum == rightSum
		if leftSum == totalSum-leftSum-val {
			return i
		}

		// Update leftSum for the next index in the loop
		leftSum += val
	}

	// Step 3: If no index satisfies the condition, return -1
	return -1
}

// Key Logic Breakdowns
// - Zero-Index Case: If the pivot is at index 0, the code correctly identifies it because leftSum starts at 0. If totalSum - nums[0] == 0, the condition is met immediately.

// - The "Total Sum" Trick: By calculating the total sum first, you avoid having to re-sum the right side of the array for every single index, which would turn an O(n) solution into a much slower O(n
// 2
//  ) solution.

// - Memory Efficiency: Go's range is used here for a clean, idiomatic iteration without needing to manage manual index counters unless necessary.

// Visualizing the State at Index 3
// For nums = [1, 7, 3, 6, 5, 6]:
// Total Sum: 28
// At Index 3 (Value 6):
// leftSum (1 + 7 + 3) = 11
// rightSum (28 - 11 - 6) = 11
// Match found! Return index 3.
