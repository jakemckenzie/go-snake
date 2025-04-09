package main

type coord struct {
	x, y int
}

const (
	Left = 1 + iota
	Right
	Up
	Down
)

const (
	HELP      = "\n\nThis game uses the W A S D keys to move the snake. \n"
	GAME_OVER = "\n\n G A M E   O V E R   ! ! ! \n"
	QUIT      = "\n\nPres the Q button the quit the game."
)
