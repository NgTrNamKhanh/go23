package ulti

//"fmt"
func stop(board [][]int, row, col int) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == 0 {
		return true
	}
	return false
}
func goRightOnly(board [][]int, row, col int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	board[row][col] = 0
	//count := 1
	goRightOnly(board, row, col+1)
	// count += goRightOnly(board, row, col+1)
	// return count
}
func goDownOnly(board [][]int, row, col int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	board[row][col] = 0
	//count := 1
	goDownOnly(board, row+1, col)
	//count += goDownOnly(board, row+1, col)
	//return count
}
func goBoth(board [][]int, row, col int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	if board[row+1][col] == 1 {
		board[row][col] = 0 // Mark cell as visited
		//count := 1
		goDownOnly(board, row+1, col)
		goBoth(board, row, col+1)
		//count += goDownOnly(board, row+1, col)
		//count += goBoth(board, row, col+1)
		//return count
	} else {
		return
		//return 0
	}
}
func rectangleHanlder(board [][]int, row, col int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	if board[row+1][col] == 0 { // go right only
		goRightOnly(board, row, col)
		//return goRightOnly(board, row, col)
	} else if board[row][col+1] == 0 { // go down only
		goDownOnly(board, row, col)
		//return goDownOnly(board, row, col)
	} else if board[row][col+1] == 0 && board[row+1][col] == 0 { //only one
		board[row][col] = 0 // Mark cell as visited
		//count := 1
		//return count
	} else { // right and down
		goBoth(board, row, col)
		//return goBoth(board, row, col)
	}
}

func CountRectangle(board [][]int) int {
	var rectangles_count int
	for row := range board {
		for col := range board {
			if board[row][col] == 1 {
				rectangles_count += 1
				rectangleHanlder(board, row, col)
				//size_of_rectangle := rectangleHanlder(board, row, col)
				//fmt.Println("Size:", size_of_rectangle)
			}
		}
	}
	return rectangles_count
}
