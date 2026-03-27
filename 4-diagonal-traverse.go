// Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.

// Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,4,7,5,3,6,8,9]

package arrayandstring

func FindDiagonalOrder(mat [][]int) []int {
	// Handle empty matrix cases
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}

	m, n := len(mat), len(mat[0])
	result := make([]int, m*n)

	// Initial position and direction
	// direction 1 means moving "up-right", -1 means "down-left"
	r, c, direction := 0, 0, 1

	for i := 0; i < m*n; i++ {
		result[i] = mat[r][c]

		if direction == 1 {
			// Moving Up-Right
			if c == n-1 {
				// Hit right wall: move down and flip direction
				r++
				direction = -1
			} else if r == 0 {
				// Hit top wall: move right and flip direction
				c++
				direction = -1
			} else {
				// Normal movement
				r--
				c++
			}
		} else {
			// Moving Down-Left
			if r == m-1 {
				// Hit bottom wall: move right and flip direction
				c++
				direction = 1
			} else if c == 0 {
				// Hit left wall: move down and flip direction
				r++
				direction = 1
			} else {
				// Normal movement
				r++
				c--
			}
		}
	}

	return result
}

// Logic Explanation

// The movement follows a zig-zag pattern. We can break down the logic into three parts:
//
// 1. The Direction FlipWe use a variable (in this case direction) to toggle between two states:
// - Up-Right: Row decreases ($r-1$), Column increases ($c+1$).
// - Down-Left: Row increases ($r+1$), Column decreases ($c-1$).

// 2. Boundary Conditions (The Order Matters!)The most critical part is handling the corners. When moving Up-Right, you might hit the top edge and the right edge simultaneously (at the last element of the first row).
// - Priority 1: If you hit the "far" boundary (Right wall for up-right, Bottom wall for down-left), you must shift to the next row/column immediately.
// - Priority 2: If you hit the "near" boundary (Top wall for up-right, Left wall for down-left), you shift to the adjacent cell.
