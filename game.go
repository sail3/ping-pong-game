package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/gookit/color"
)

const (
	BallRadius = 0
	PlayerSize = 3
)

type (
	cell struct {
		isWall  bool
		visited bool
	}

	point struct {
		X int
		Y int
	}

	cells [][]cell

	board struct {
		cells     cells
		playerA   point
		playerB   point
		turn      string
		hasWinner bool
		ball      point
		incX      int
		incY      int
	}
)

func (b *board) HandlePlayerA(r rune) {
	switch {
	case r == 'w':
		if b.playerA.Y >= 0 {
			b.playerA.Y--
		}
	case r == 's':
		if b.playerA.Y < len(b.cells) {
			b.playerA.Y++
		}
	}
}
func (b *board) HandlePlayerB(key keyboard.Key) {
	switch {
	case key == keyboard.KeyArrowUp:
		if b.playerB.Y >= 0 {
			b.playerB.Y--
		}
	case key == keyboard.KeyArrowDown:
		if b.playerB.Y < len(b.cells) {
			b.playerB.Y++
		}
	}
}

func (b board) String() string {
	var s string
	for i, row := range b.cells {
		for j, col := range row {
			switch {
			case b.playerA.X == j && (b.playerA.Y-PlayerSize < i && b.playerA.Y+PlayerSize > i):
				s += color.Red.Sprintf("▒")
			case b.playerB.X == j && (b.playerB.Y-PlayerSize < i && b.playerB.Y+PlayerSize > i):
				s += color.Blue.Sprintf("▒")
			case (b.ball.X-BallRadius <= j && b.ball.X+BallRadius >= j) && (b.ball.Y-BallRadius <= i && b.ball.Y+BallRadius >= i) ||
				(b.ball.X-BallRadius+1 <= j && b.ball.X+BallRadius+1 >= j) && (b.ball.Y-BallRadius <= i && b.ball.Y+BallRadius >= i):
				s += color.Black.Sprintf("▒")
			case i == 0 || i == len(b.cells)-1 || j == len(row)/2:
				s += color.White.Sprintf("▒")
			case col.isWall:
				s += color.LightGreen.Sprintf("▒")
			default:
				s += " "
			}
		}
		s += fmt.Sprintln()
	}
	return s
}

func (b *board) MoveBall() {
	if b.ball.Y+b.incY >= len(b.cells)-1 || b.ball.Y+b.incY < 1 {
		b.incY *= -1
	}
	if b.ball.X+b.incX+1 == b.playerB.X &&
		(b.ball.Y+b.incY < b.playerB.Y+PlayerSize && b.ball.Y+b.incY > b.playerB.Y-PlayerSize) {
		b.incX *= -1
		b.turn = "Player02"
	}
	if b.ball.X+b.incX == b.playerA.X &&
		(b.ball.Y+b.incY < b.playerA.Y+PlayerSize && b.ball.Y+b.incY > b.playerA.Y-PlayerSize) {
		b.incX *= -1
		b.turn = "Player01"
	}
	if b.ball.X+b.incX+1 >= len(b.cells[0]) || b.ball.X+b.incX < 0 {
		b.incX *= -1
		b.hasWinner = true
	}

	b.ball.X += b.incX
	b.ball.Y += b.incY
}
