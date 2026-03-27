// Given an integer numRows, return the first numRows of Pascal's triangle.
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

package arrayandstring

func Generate(numRows int) [][]int {
	// Initialize the result slice with the number of rows requested
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		// Create a row with i+1 elements (Row 0 has 1 element, Row 1 has 2, etc.)
		row := make([]int, i+1)

		// Every row starts and ends with 1
		row[0], row[i] = 1, 1

		// Fill in the interior elements
		// These are the sum of the two elements above it from the previous row
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result[i] = row
	}

	return result
}

// Logic Explanation

// To build the triangle efficiently, we use the properties of the rows and the relationship with the "parent" row:
// - Row Definition: The $n$-th row (starting from index 0) always has exactly $n+1$ elements.
// - The "1" Bookends: The first and last elements of any row are always 1. We set these manually.
// - The Summation Rule: For any element at index $j$ (where $j$ is not the first or last element), its value is:

// result[i][j] = result[i-1][j-1] + result[i-1][j]

// This means we simply look at the previous row (i-1) and add the element at the same column index (j) and the one just to the left of it (j-1).

// Complexity
// Time Complexity: $O(numRows^2)$. We use a nested loop to fill each element. The total number of elements is $\frac{n(n+1)}{2}$.
// Space Complexity: $O(numRows^2)$ to store the triangle. If we only consider extra auxiliary space used during calculation, it is $O(1)$ beyond the output.

// Quick Tip for GoNotice that we use make([]int, i+1) inside the loop. This is more efficient than using append repeatedly because we already know exactly how long each row needs to be.
