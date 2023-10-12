package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		log.Fatal("unsupported platform. cannot clear console.")
	}
}

func initializeBoard() *board {
	var b *board
	b = &board{
		cells: initializeCells(30, 10),
		playerA: point{
			X: 0,
			Y: 10,
		},
		playerB: point{
			X: 60,
			Y: 10,
		},
		ball: point{
			X: 50,
			Y: 0,
		},
		incX: 1,
		incY: 1,
	}
	return b
}

func initializeCells(width, height int) cells {
	cells := make(cells, height*2+1)
	for i := range cells {
		cells[i] = make([]cell, width*2+1)
		for j := range cells[i] {
			cells[i][j].isWall = true
		}
	}
	return cells
}
