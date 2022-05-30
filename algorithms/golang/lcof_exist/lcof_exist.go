package lcof_exist

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if search(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func search(board [][]byte, word string, i, j, k int) bool {
	if i > len(board)-1 || i < 0 || j > len(board[i])-1 || j < 0 || board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	tmp := board[i][j]
	board[i][j] = ' '

	match := search(board, word, i+1, j, k+1) || search(board, word, i, j+1, k+1) || search(board, word, i-1, j, k+1) || search(board, word, i, j-1, k+1)

	board[i][j] = tmp

	return match
}
