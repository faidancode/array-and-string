// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).

package arrayandstring

func RemoveElement(nums []int, val int) int {
	left := 0
	right := len(nums)

	for left < right {
		if nums[left] == val {
			// "Throw" the unwanted value to the end by
			// replacing it with the last available element.
			nums[left] = nums[right-1]
			right--
			// We don't increment 'left' here because the new element
			// at nums[left] might also be equal to 'val'.
		} else {
			left++
		}
	}
	return left
}
