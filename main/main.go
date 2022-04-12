package main

import (
	"fmt"
)

func main() {
	//showBoard(solution(3))
	//showBoard(solution(5))
	showBoard(solution(7))
}

func showBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
}

func makeBoard(number int) [][]int {
	board := make([][]int, number)

	for i := 0; i < number; i++ {
		board[i] = make([]int, number)
	}
	return board
}

func solution(number int) [][]int {
	board := makeBoard(number)
	if number%2 == 0 {
		return evenSolve(board, number, number/2, 0, 1)
	}
	return oddSolve(board, number, number/2, 0, 1, 0)
}

func oddSolve(board [][]int, number int, dx int, dy int, count int, chagneNum int) [][]int {
	if chagneNum == number*number {
		return board
	}
	board[dy][dx] = count
	chagneNum++
	count++
	changeDx, changeDy := oddSwitch(dx+1, dy-1, number)
	if board[changeDy][changeDx] == 0 {
		return oddSolve(board, number, changeDx, changeDy, count, chagneNum)
	}
	dy = dy + 1
	return oddSolve(board, number, dx, dy, count, chagneNum)
}

func oddSwitch(dx int, dy int, number int) (int, int) {
	switch {
	case dx < 0:
		dx = number + dx
	case dx >= number:
		dx = dx - number
	}

	switch {
	case dy < 0:
		dy = number + dy
	case dy >= number:
		dy = dy - number
	}

	return dx, dy
}

func evenSolve(board [][]int, number int, dx int, dy int, count int) [][]int {
	return board
}
