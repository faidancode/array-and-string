// Given two binary strings a and b, return their sum as a binary string.
// Input: a = "11", b = "1"
// Output: "100"

package arrayandstring

import (
	"strconv"
)

func addBinary(a string, b string) string {
	res := ""
	i, j := len(a)-1, len(b)-1
	carry := 0

	// Keep going as long as there is a bit to process or a carry left
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			// Convert character '0'/'1' to integer 0/1
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// The current bit is sum % 2 (either 0 or 1)
		res = strconv.Itoa(sum%2) + res
		// The carry is sum / 2 (1 if sum was 2 or 3)
		carry = sum / 2
	}

	return res
}

// The Manual Strategy

// Pointers: Initialize i at the end of string a and j at the end of string b.

// Carry: Use an integer carry initialized to 0.

// Loop: While there are characters left in a OR b, or if there is a remaining carry:
// - Add the bit from a (if i >= 0) to sum.
// - Add the bit from b (if j >= 0) to sum.
// - The new bit for the result is sum % 2.
// - The new carry is sum / 2.

// Reverse: Since we build the result from right to left, we need to reverse the final string.
