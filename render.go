package main

import "fmt"

func RenderTitle() string {
	return "Initial attempt at Snake game in Go making sure I'm using slices/pointers/channels properly:\n"
}

func RenderArena(m *Model) {
	m.arena = make([][]string, m.height)
	for i := range m.arena {
		m.arena[i] = make([]string, m.width)
		for j := range m.arena[i] {
			if i == 0 || i == m.height-1 || j == 0 || j == m.width-1 {
				m.arena[i][j] = m.horizontalLine
			} else {
				m.arena[i][j] = m.emptySymbol
			}
		}
	}
}

func RenderSnake(m *Model) {
	for _, c := range m.snake.body {
		m.arena[c.x][c.y] = m.snakeSymbol
	}
}

func RenderEgg(m *Model) {
	if m.egg.x >= 0 && m.egg.x < m.height && m.egg.y >= 0 && m.egg.y < m.width {
		m.arena[m.egg.x][m.egg.y] = m.eggSymbol
	} else {
		fmt.Printf("Invalid egg position: x=%d, y=%d\n", m.egg.x, m.egg.y)
	}
}

func RenderScore(score int) string {
	return fmt.Sprintf("Score: %d", score)
}

func RenderGameOver() string {
	return GAME_OVER
}

func RenderHelp(text string) string {
	return text
}
