// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// Input: s = "the sky is blue"
// Output: "blue is sky the"

package arrayandstring

// Solution for Reverse Words in a String (Zero Imports)
// File: leetcode/reverse_words.go

func reverseWords(s string) string {
	// Step 1: Convert to byte slice and clean up spaces
	b := []byte(s)
	n := len(b)

	// Use two pointers to remove extra spaces in-place
	// 	- Fast Pointer (The Scout): This pointer scans the entire array, looking for "good" data (like a word) and skipping "bad" data (like extra spaces).
	// - Slow Pointer (The Worker): This pointer stays behind and marks the position where the next piece of "good" data should be written.

	slow, fast := 0, 0
	for fast < n {
		// Skip leading/multiple spaces
		for fast < n && b[fast] == ' ' {
			fast++
		}
		if fast == n {
			break
		}
		// Add a single space before a word (except for the first word)
		if slow != 0 {
			b[slow] = ' '
			slow++
		}
		// Copy the word
		for fast < n && b[fast] != ' ' {
			b[slow] = b[fast]
			slow++
			fast++
		}
	}
	// Trim the slice to the actual content length
	b = b[:slow]
	n = len(b)

	// Helper to reverse a segment
	reverse := func(data []byte, start, end int) {
		for start < end {
			data[start], data[end] = data[end], data[start]
			start++
			end--
		}
	}

	// Step 2: Reverse the entire cleaned slice
	// "the sky" -> "yks eht"
	reverse(b, 0, n-1)

	// Step 3: Reverse each individual word
	// "yks eht" -> "sky the"
	wordStart := 0
	for i := 0; i <= n; i++ {
		if i == n || b[i] == ' ' {
			reverse(b, wordStart, i-1)
			wordStart = i + 1
		}
	}

	return string(b)
}
