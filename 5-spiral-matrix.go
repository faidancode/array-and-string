package arrayandstring

func SpiralOrder(matrix [][]int) []int {
	// Base case: check for empty matrix or empty rows
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// Initialize boundaries
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	result := []int{}

	// Continue as long as the boundaries haven't crossed each other
	for top <= bottom && left <= right {

		// 1. Move Right: Traverse the current top row from left to right
		for c := left; c <= right; c++ {
			result = append(result, matrix[top][c])
		}
		top++ // Move the ceiling down

		// 2. Move Down: Traverse the current right column from top to bottom
		for r := top; r <= bottom; r++ {
			result = append(result, matrix[r][right])
		}
		right-- // Move the right wall in

		// 3. Move Left: Only if there is still a row to traverse
		// This prevents double-processing in 1-row matrices
		if top <= bottom {
			for c := right; c >= left; c-- {
				result = append(result, matrix[bottom][c])
			}
			bottom-- // Move the floor up
		}

		// 4. Move Up: Only if there is still a column to traverse
		// This prevents double-processing in 1-column matrices
		if left <= right {
			for r := bottom; r >= top; r-- {
				result = append(result, matrix[r][left])
			}
			left++ // Move the left wall in
		}
	}

	return result
}

// Why the "If" Checks are EssentialThe spiral logic effectively "eats" the matrix from the outside in. Here is why the logic requires those specific guards:
// 1. Shrinking Boundaries: Every time we finish a side, we move a boundary (top++, right--, etc.).
// 2. The "Single Row" Scenario: Imagine a $1 \times 3$ matrix.
// - Step 1 (Right): You traverse all 3 elements. top becomes 1.
// - Step 2 (Down): The loop for r := top; r <= bottom won't run because $1$ is not $\le 0$. right becomes 1.
// -Step 3 (Left): Without the if top <= bottom check, the code would try to run a loop on matrix[bottom] (index 0), but it would be moving backward. Since right is now 1 and left is 0, it would incorrectly add elements you already visited.
// 3.Safety: The if statements act as a mid-loop "emergency brake" to ensure that the boundary we just moved didn't just end the entire process.
