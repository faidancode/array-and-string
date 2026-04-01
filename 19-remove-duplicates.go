package arrayandstring

func removeDuplicates(nums []int) int {
	// Handle the edge case for an empty slice
	if len(nums) == 0 {
		return 0
	}

	// 'i' is the slow pointer that tracks the position of the last unique element
	i := 0

	// 'j' is the fast pointer that explores the slice
	for j := 1; j < len(nums); j++ {
		// If we find a value that is different from our last unique element...
		if nums[j] != nums[i] {
			// ...move the slow pointer forward and update it with the new unique value
			i++
			nums[i] = nums[j]
		}
	}

	// The number of unique elements is the index 'i' plus 1 (since it's 0-indexed)
	return i + 1
}
