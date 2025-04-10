package main

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

type Model struct {
	horizontalLine, emptySymbol, snakeHeadSymbol, snakeBodySymbol, eggSymbol string
	width, height                                                            int
	arena                                                                    [][]string
	snake                                                                    snake
	lostGame, newHighScore                                                   bool
	score, highScore                                                         int
	egg                                                                      coord
	rng                                                                      *rand.Rand
	directionQueue                                                           []int
}

func initialModel() *Model {
	return &Model{
		horizontalLine:  "🧱",
		emptySymbol:     "  ",
		snakeHeadSymbol: "🟢",
		snakeBodySymbol: "🟩",
		eggSymbol:       "🥚",
		width:           40,
		height:          20,
		arena:           [][]string{},
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
		lostGame:     false,
		score:        0,
		egg:          coord{x: 10, y: 10},
		rng:          rand.New(rand.NewSource(time.Now().UnixNano())),
		highScore:    readHighScore(),
		newHighScore: false,
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

func (m *Model) isOppositeDirection(dir1, dir2 int) bool {
	opposites := map[int]int{
		Up:    Down,
		Down:  Up,
		Left:  Right,
		Right: Left,
	}
	return opposites[dir1] == dir2
}

func (m *Model) changeSnakeDirection(direction int) tea.Cmd {
	if len(m.directionQueue) == 0 {
		if !m.isOppositeDirection(direction, m.snake.direction) {
			m.directionQueue = append(m.directionQueue, direction)
		}
	} else {
		lastQueued := m.directionQueue[len(m.directionQueue)-1]
		if !m.isOppositeDirection(direction, lastQueued) {
			m.directionQueue = append(m.directionQueue, direction)
		}
	}

	if m.snake.hitWall(*m) {
		m.lostGame = true
		if m.score > m.highScore {
			m.newHighScore = true
			m.highScore = m.score
			writeHighScore(m.score)
		}
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
	if len(m.directionQueue) > 0 {
		newDir := m.directionQueue[0]
		m.directionQueue = m.directionQueue[1:]
		if !m.isOppositeDirection(newDir, m.snake.direction) {
			m.snake.direction = newDir
		}
	}

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
