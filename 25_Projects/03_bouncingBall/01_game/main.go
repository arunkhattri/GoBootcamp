package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width     = 50
		height    = 10
		cellBall  = '⚽'
		cellEmpty = ' '
		maxFrames = 60 * 20
		speed     = time.Second / 10
	)
	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)
	board := make([][]bool, width)

	for column := range board {
		board[column] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		//remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}
		// ball position
		board[px][py] = true

		// buffer - to draw board and printing once in each loop
		buf := make([]rune, width*height)
		buf = buf[:0] // reducing the length to 0 for reuse of backing array

		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				// fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
