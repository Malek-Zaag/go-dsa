package backtracking

const (
	empty = iota
	queen
)

type chessboard [][]int

// NQueens returns possible solutions to the n-queen puzzle in an n x n chessboard
// where n queens are placed on the chessboard such that none attacks another
func NQueens(n int) []chessboard {
	output := nQueensRecursive(0, n, make([]int, n), make(chessboard, 0))
	return toPrettyChessboard(output, n)
}

func nQueensRecursive(row, n int, cols []int, output chessboard) chessboard {
	if n == 0 {
		return output
	}
	if row == n {
		output = append(output, append([]int{}, cols...))
		return output
	}
	for col := 0; col < n; col++ {
		if isValidQueenPlacement(row, col, cols) {
			cols[row] = col
			output = nQueensRecursive(row+1, n, cols, output)
		}
	}
	return output
}

func isValidQueenPlacement(row, col int, cols []int) bool {
	for i := 0; i < row; i++ {
		if col == cols[i] || col-row == cols[i]-i || col+row == cols[i]+i {
			return false
		}
	}
	return true
}

func toPrettyChessboard(solutions chessboard, n int) []chessboard {
	chessboards := make([]chessboard, len(solutions))
	for i, row := range solutions {
		chessBoard := make(chessboard, n)
		for j, col := range row {
			newRow := make([]int, n)
			newRow[col] = queen
			chessBoard[j] = newRow
		}
		chessboards[i] = chessBoard
	}
	return chessboards
}
