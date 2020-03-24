package main

import "fmt"

/*
On an 8 x 8 chessboard, there is one white rook.  There also may be empty squares, white bishops, and black pawns.
These are given as characters 'R', '.', 'B', and 'p' respectively.
Uppercase characters represent white pieces, and lowercase characters represent black pieces.
The rook moves as in the rules of Chess: it chooses one of four cardinal directions (north, east, west, and south),
then moves in that direction until it chooses to stop, reaches the edge of the board,
or captures an opposite colored pawn by moving to the same square it occupies.
Also, rooks cannot move into the same square as other friendly bishops.
Return the number of pawns the rook can capture in one move.

Example 1:
Input: [[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".","R",".",".",".","p"],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 3
Explanation:
In this example the rook is able to capture all the pawns.

Example 2:
Input: [[".",".",".",".",".",".",".","."],
		[".","p","p","p","p","p",".","."],
		[".","p","p","B","p","p",".","."],
		[".","p","B","R","B","p",".","."],
		[".","p","p","B","p","p",".","."],
		[".","p","p","p","p","p",".","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 0
Explanation:
Bishops are blocking the rook to capture any pawn.

Example 3:
Input: [[".",".",".",".",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		["p","p",".","R",".","p","B","."],
		[".",".",".",".",".",".",".","."],
		[".",".",".","B",".",".",".","."],
		[".",".",".","p",".",".",".","."],
		[".",".",".",".",".",".",".","."]]
Output: 3
Explanation:
The rook can capture the pawns at positions b5, d6 and f5.

Note:
board.length == board[i].length == 8
board[i][j] is either 'R', '.', 'B', or 'p'
There is exactly one cell with board[i][j] == 'R'
*/

func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}

func numRookCaptures(board [][]byte) int {
	pRowM := make(map[int]int, len(board))
	pColM := make(map[int]int, len(board[0]))
	bRowM := make(map[int]int, len(board))
	bColM := make(map[int]int, len(board[0]))
	cnt := 0
	findR := false
	rI := -1
	rJ := -1

	for i := 0; i < len(board) && !findR; i++ {
		for j := 0; j < len(board[0]) && !findR; j++ {
			if board[i][j] == 'R' {
				findR = true
				rI = i
				rJ = j
				break
			} else if board[i][j] == 'p' {
				pRowM[i] = j
				pColM[j] = i
			} else if board[i][j] == 'B' {
				bRowM[i] = j
				bColM[j] = i
			}
		}
	}

	if p, ok := pRowM[rI]; ok {
		if b, bOk := bRowM[rI]; bOk {
			if b < p {
				cnt++
			}
		} else {
			cnt++
		}
	}

	if p, ok := pColM[rJ]; ok {
		if b, bOk := bColM[rJ]; bOk {
			if b < p {
				cnt++
			}
		} else {
			cnt++
		}
	}

	for j := rJ + 1; j < len(board[0]); j++ {
		if board[rI][j] == 'B' {
			break
		}

		if board[rI][j] == 'p' {
			cnt++
			break
		}
	}

	for i := rI + 1; i < len(board); i++ {
		if board[i][rJ] == 'B' {
			break
		}

		if board[i][rJ] == 'p' {
			cnt++
			break
		}
	}

	return cnt
}
