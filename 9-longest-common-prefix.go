// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

package arrayandstring

import "strings"

func LongestCommonPrefix(strs []string) string {
	// If the input slice is empty, there is no common prefix
	if len(strs) == 0 {
		return ""
	}

	// Initialize the prefix with the first string in the slice
	prefix := strs[0]

	// Iterate through the rest of the strings in the slice
	for i := 1; i < len(strs); i++ {
		// While the current string does not start with the current prefix
		for !strings.HasPrefix(strs[i], prefix) {
			// If the prefix becomes empty, it means there is no common prefix at all
			if len(prefix) == 0 {
				return ""
			}

			// Reduce the length of the prefix by one character from the end
			// This is the slicing operation that removes the last character
			prefix = prefix[:len(prefix)-1]
		}
	}

	// Return the longest common prefix found among all strings
	return prefix
}

// To find the longest common prefix among an array of strings, the most efficient approach is "Horizontal Scanning." We assume the first string is the prefix and then compare it against each subsequent string, shortening it whenever a mismatch is found.

// Logic Breakdown

// Edge Case: If the input array is empty, return an empty string.

// Initialization: Start by assuming the first string (strs[0]) is the prefix.

// Iteration: Loop through the rest of the strings in the array.

// Comparison: For each string, check if it starts with the current prefix.
// - If it doesn't, remove the last character from the prefix and try again.
// - If the prefix becomes empty during this process, there is no common prefix, so return "".

// Return: After checking all strings, the remaining prefix is the longest common one.
