package ulti

//"fmt"
func stop(board [][]int, row, col int) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] == 0 {
		return true
	}
	return false
}
func goRightOnly(board [][]int, row, col int, flag bool) int {
	if stop(board, row, col) {
		return col -1
		//return 0
	}
	if !flag{
		board[row][col] = 0
		//count := 1
		goRightOnly(board, row, col+1, flag)
		return 0
		// count += goRightOnly(board, row, col+1)
		// return count
	}else{
		if board[row+1][col] == 1 {
			board[row][col] = 0 // Mark cell as visited
			maxCol:= goRightOnly(board, row, col+1, flag)
			return maxCol
		}else{
			return col -1
		}
	}
	
}
func goDownOnly(board [][]int, row, col, maxCol int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	if maxCol == 0 {
		board[row][col] = 0
		//count := 1
		goDownOnly(board, row+1, col, maxCol)
		//count += goDownOnly(board, row+1, col)
		//return count
	} else {
		var valid bool
		for i := col; i <= maxCol; i++ {
			if board[row][i] == 1 {
				valid = true
			} else {
				valid = false
			}
		}
		if valid {
			for i := col; i <= maxCol; i++ {
				board[row][i] = 0
			}
			goDownOnly(board, row+1, col, maxCol)
		} else {
			return
		}

	}

}
func goBoth(board [][]int, row, col int) {
	if stop(board, row, col) {
		return 
		//return 0
	}
	board[row][col] = 0 // Mark cell as visited
	//count := 1
	flag := true
	maxCol := goRightOnly(board, row, col+1, flag)
	goDownOnly(board, row+1, col, maxCol)
	return
	//count += goDownOnly(board, row+1, col)
	//count += goBoth(board, row, col+1)
	//return count
}
func rectangleHanlder(board [][]int, row, col int) {
	if stop(board, row, col) {
		return
		//return 0
	}
	if stop(board, row+1, col) { // go right only
		flag:= false
		goRightOnly(board, row, col, flag)
		//return goRightOnly(board, row, col)
	} else if stop(board, row, col+1) { // go down only
		goDownOnly(board, row, col, 0)
		//return goDownOnly(board, row, col)
	} else if stop(board, row, col+1) && stop(board, row+1, col) { //only one
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
