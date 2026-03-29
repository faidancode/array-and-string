package arrayandstring

// ReverseString reverses a slice of characters in-place.
func ReverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		// Swap the elements
		s[left], s[right] = s[right], s[left]

		// Move pointers toward the middle
		left++
		right--
	}
}

// Key Details
// Time Complexity: $O(n)$, where $n$ is the length of the string. We visit each element once.
// Space Complexity: $O(1)$, as we only use two integer variables (left and right) regardless of the input size.
// Go Syntax: We use the tuple assignment a, b = b, a to perform the swap in a single line, which is idiomatic in Go and avoids the need for an explicit temp variable.
