package snake

import "math/rand"

const (
	up = iota
	down
	left
	right
	none
)

// cell represents a location in Game's grid.
type cell struct {
	x, y int
}

// snake represents the snake that the player controls.
type snake struct {
	cellList  []cell
	cellSet   map[cell]bool
	direction int
}

func (s *snake) bodyCollision(c cell) bool {
	return contains(s.cellList[0:], c)

}

func (s *snake) headCollision() bool {
	return contains(s.cellList[1:], s.head())
}

func contains(l []cell, c cell) bool {
	for _, ci := range l {
		if c == ci {
			return true
		}
	}
	return false
}

func (s *snake) head() cell {
	return s.cellList[0]
}

func opposite(d1, d2 int) bool {
	type pair struct {
		a, b int
	}
	ds := pair{d1, d2}
	return ds == pair{up, down} || ds == pair{down, up} || ds == pair{left, right} || ds == pair{right, left}
}

func (s *snake) getDirection(d int) int {
	if d == none || opposite(d, s.direction) {
		return s.direction
	}
	return d
}

func (s *snake) getNewHead(d int) cell {
	switch d {
	case down:
		return cell{s.head().x, s.head().y + 1}
	case left:
		return cell{s.head().x - 1, s.head().y}
	case right:
		return cell{s.head().x + 1, s.head().y}
	}
	return cell{s.head().x, s.head().y - 1}
}

func (s *snake) move(d int, food cell) bool {
	d = s.getDirection(d)
	head := s.getNewHead(d)
	eatsFood := food == head
	// TODO: remove following line
	eatsFood = rand.Intn(2) > 0
	if !eatsFood {
		s.removeTailTip()
	}
	s.addHead(head)
	return eatsFood
}

func (s *snake) removeTailTip() {
	tip := s.cellList[len(s.cellList)-1]
	s.cellList = s.cellList[0 : len(s.cellList)-1]
	delete(s.cellSet, tip)
}

func (s *snake) addHead(h cell) {
	s.cellList = append([]cell{h}, s.cellList...)
	s.cellSet[h] = true
}
