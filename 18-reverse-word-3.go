// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Input: s = "Mr Ding"
// Output: "rM gniD"

package arrayandstring

func reverseWords3(s string) string {
	// Convert to byte slice for in-place modification
	b := []byte(s)
	n := len(b)
	slow := 0

	for fast := 0; fast <= n; fast++ {
		// We found a space or reached the very end of the string
		if fast == n || b[fast] == ' ' {
			// Reverse the word found between slow and fast-1
			left := slow
			right := fast - 1

			for left < right {
				b[left], b[right] = b[right], b[left]
				left++
				right--
			}

			// Move slow pointer to the beginning of the next word
			slow = fast + 1
		}
	}

	return string(b)
}

// Visualization of the Pointers

// Initial: slow and fast are at index 0 ('L').
// Scanning: fast moves until it hits the first space after "Let's".
// Reversing: The reverse logic swaps 'L' with 's', then 'e' with 't', etc., resulting in "s'teL".
// Skipping: fast is at the space. slow jumps to fast + 1 (the 't' in "take").
// Repeat: This continues until the end of the string.
