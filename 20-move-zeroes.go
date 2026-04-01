// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]

package arrayandstring

func moveZeroes(nums []int) {
	// lastNonZeroIndex keeps track of the position where the
	// next non-zero element should be moved.
	lastNonZeroIndex := 0

	// Iterate through the array with the 'cur' pointer
	for cur := 0; cur < len(nums); cur++ {
		// If the current element is not zero
		if nums[cur] != 0 {
			// Swap the elements at 'lastNonZeroIndex' and 'cur'
			// Using Go's tuple assignment for a clean swap
			nums[lastNonZeroIndex], nums[cur] = nums[cur], nums[lastNonZeroIndex]

			// Move the pointer forward for the next non-zero element
			lastNonZeroIndex++
		}
	}
}

// The Logic
// 1. First Pointer (lastNonZeroIndex): This keeps track of where the next non-zero element should be placed.

// 2. Second Pointer (cur): This iterates through the entire array.

// 3. Whenever the cur pointer encounters a non-zero element, we swap it with the element at lastNonZeroIndex and increment lastNonZeroIndex.
