package main

type TicTacToe struct {
	cols   []int
	rows   []int
	diags  []int
	rdiags []int
	n      int
}

func Constructor348(n int) TicTacToe {
	ttt := TicTacToe{}
	ttt.cols = make([]int, n)
	ttt.rows = make([]int, n)
	ttt.diags = make([]int, 2*n-1)
	ttt.rdiags = make([]int, 2*n-1)
	ttt.n = n
	return ttt
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
//lint:ignore ST1006 this
func (this *TicTacToe) Move(row int, col int, player int) int {
	d := 1
	if player == 2 {
		d = -1
	}

	this.rows[row] += d
	this.cols[col] += d
	this.diags[row+col] += d
	this.rdiags[row-col+this.n-1] += d
	if this.rows[row] == this.n || this.cols[col] == this.n || this.diags[row+col] == this.n || this.rdiags[row-col+this.n-1] == this.n {
		return 1
	}
	if this.rows[row] == -this.n || this.cols[col] == -this.n || this.diags[row+col] == -this.n || this.rdiags[row-col+this.n-1] == -this.n {
		return 2
	}

	return 0
}
