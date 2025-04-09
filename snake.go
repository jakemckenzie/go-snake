package main

type snake struct {
	body      []coord
	length    int
	direction int
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
