package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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

type TickMsg time.Time

type Model struct {
	horizontalLine string
	verticalLine   string
	emptySymbol    string
	snakeSymbol    string
	eggSymbol      string
	width          int
	height         int
	arena          [][]string
	snake          snake
	lostGame       bool
	score          int
	egg            egg
	rng            *rand.Rand
}

type snake struct {
	body      []coord
	length    int
	direction int
}

type egg struct {
	x, y int
}

type coord struct {
	x, y int
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been some errors initializing the game mode: %v", err)
		os.Exit(1)
	}
}

func initialModel() *Model {
	return &Model{
		horizontalLine: "🧱",
		verticalLine:   "🧱",
		emptySymbol:    "  ",
		snakeSymbol:    "🟩",
		eggSymbol:      "🥚",
		width:          40,
		height:         20,
		arena:          [][]string{},
		snake: snake{
			body: []coord{
				{x: 11, y: 11},
				{x: 12, y: 12},
				{x: 13, y: 13},
				{x: 14, y: 14},
			},
			length:    3,
			direction: Right,
		},
		lostGame: false,
		score:    0,
		egg: egg{
			x: 10, y: 10,
		},
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m *Model) Init() tea.Cmd {
	var x, y int
	x = rand.Intn(m.height-2) + 1
	y = rand.Intn(m.width-2) + 1
	m.egg.x = x
	m.egg.y = y
	return m.tick()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl + c", "q":
			return m, tea.Quit
		case "up", "w":
			return m, m.changeSnakeDirection(Up)
		case "down", "s":
			return m, m.changeSnakeDirection(Down)
		case "left", "a":
			return m, m.changeSnakeDirection(Left)
		case "right", "d":
			return m, m.changeSnakeDirection(Right)
		}
	case TickMsg:
		return m, m.moveSnake()
	}
	return m, nil
}

func (m *Model) View() string {
	var sb strings.Builder

	sb.WriteString(RenderTitle())
	sb.WriteByte('\n')

	RenderArena(m)
	RenderSnake(m)
	RenderEgg(m)

	for _, row := range m.arena {
		sb.WriteString(strings.Join(row, "") + "\n")
	}

	sb.WriteByte('\n')
	sb.WriteString(RenderScore(m.score))
	sb.WriteByte('\n')

	if m.lostGame {
		sb.WriteString(RenderGameOver())
	}

	sb.WriteString(RenderHelp(HELP))
	sb.WriteByte('\n')
	sb.WriteString(RenderHelp("Press q or ctrl+c to quit."))
	sb.WriteByte('\n')
	sb.WriteByte('\n')

	return sb.String()
}

func (m *Model) tick() tea.Cmd {
	return tea.Tick(time.Second/10, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m *Model) changeSnakeDirection(direction int) tea.Cmd {
	if m.snake.hitWall(*m) {
		m.lostGame = true
		return tea.Quit
	}

	opposites := map[int]int{
		Up:    Down,
		Down:  Up,
		Left:  Right,
		Right: Left,
	}

	if opposites[direction] != m.snake.direction {
		m.snake.direction = direction
	}

	return nil
}

func (m *Model) moveSnake() tea.Cmd {
	h := m.snake.getHead()
	c := coord{x: h.x, y: h.y}

	switch m.snake.direction {
	case Right:
		c.y++
	case Left:
		c.y--
	case Up:
		c.x--
	case Down:
		c.x++
	}

	if c.x == m.egg.x && c.y == m.egg.y {
		m.snake.length++
		m.snake.body = append(m.snake.body, c)
		for {
			x := rand.Intn(m.height-2) + 1
			y := rand.Intn(m.width-2) + 1
			if !m.snake.hitSelf(coord{x, y}) {
				m.egg.x = x
				m.egg.y = y
				break
			}
		}
		m.score += 10
	} else if m.snake.hitWall(*m) || m.snake.hitSelf(c) {
		m.lostGame = true
		return tea.Quit
	} else {
		m.snake.body = append(m.snake.body[1:], c)
	}

	return m.tick()
}

func (s *snake) hitWall(m Model) bool {
	h := s.getHead()
	return h.x <= 0 || h.x >= m.height-1 || h.y <= 0 || h.y >= m.width-1
}

func (s *snake) hitSelf(c coord) bool {
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}
	return false
}

func (s *snake) getHead() coord {
	return s.body[len(s.body)-1]
}

func RenderTitle() string {
	return "Initial attempt at Snake Game in Go using slices/pointers/channels:\n"
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
