// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

// Input: rowIndex = 3
// Output: [1,3,3,1]

// Input: rowIndex = 1
// Output: [1,1]

// 1
// 1 1
// 1 2 1
// 1 3 3 1

package arrayandstring

func getRow(rowIndex int) []int {
	// Initialize the row with size rowIndex + 1
	row := make([]int, rowIndex+1)

	// The first element of every row is always 1
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		// Update from right to left to avoid using values
		// that have already been updated for the current row.
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}
