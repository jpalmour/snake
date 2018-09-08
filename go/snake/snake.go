package snake

import (
	"github.com/jpalmour/snake/go"
)

// Snake represents the snake that the player controls.
type Snake struct {
	cellList   []snakeapp.Cell
	cellSet    map[snakeapp.Cell]bool
	direction  int
	controller snakeapp.Controller
}

func New(boardSize int, c snakeapp.Controller) *Snake {
	s := &Snake{
		cellList:   []snakeapp.Cell{},
		cellSet:    map[snakeapp.Cell]bool{},
		direction:  snakeapp.Up,
		controller: c,
	}
	s.addHead(snakeapp.Cell{boardSize / 2, boardSize / 2})
	return s
}

func (s *Snake) CellSet() map[snakeapp.Cell]bool {
	return s.cellSet
}

func (s *Snake) CellList() []snakeapp.Cell {
	return s.cellList
}

func (s *Snake) BodyCollision(c snakeapp.Cell) bool {
	return contains(s.cellList[0:], c)
}

func (s *Snake) HeadCollision() bool {
	return contains(s.cellList[1:], s.Head())
}

func (s *Snake) Head() snakeapp.Cell {
	return s.cellList[0]
}

func (s *Snake) Move(food snakeapp.Cell) bool {
	s.direction = s.getDirection()
	head := s.getNewHead(s.direction)
	eatsFood := food == head
	if !eatsFood {
		s.removeTailTip()
	}
	s.addHead(head)
	return eatsFood
}

func contains(l []snakeapp.Cell, c snakeapp.Cell) bool {
	for _, ci := range l {
		if c == ci {
			return true
		}
	}
	return false
}

func opposite(d1, d2 int) bool {
	type pair struct {
		a, b int
	}
	ds := pair{d1, d2}
	return ds == pair{snakeapp.Up, snakeapp.Down} || ds == pair{snakeapp.Down, snakeapp.Up} || ds == pair{snakeapp.Left, snakeapp.Right} || ds == pair{snakeapp.Right, snakeapp.Left}
}

func (s *Snake) getDirection() int {
	d := s.controller.GetDirection()
	if opposite(d, s.direction) {
		return s.direction
	}
	return d
}

func (s *Snake) getNewHead(d int) snakeapp.Cell {
	switch d {
	case snakeapp.Up:
		return snakeapp.Cell{s.Head().X, s.Head().Y - 1}
	case snakeapp.Down:
		return snakeapp.Cell{s.Head().X, s.Head().Y + 1}
	case snakeapp.Left:
		return snakeapp.Cell{s.Head().X - 1, s.Head().Y}
	case snakeapp.Right:
		return snakeapp.Cell{s.Head().X + 1, s.Head().Y}
	}
	return snakeapp.Cell{s.Head().X + 1, s.Head().Y}
}

func (s *Snake) removeTailTip() {
	tip := s.cellList[len(s.cellList)-1]
	s.cellList = s.cellList[0 : len(s.cellList)-1]
	delete(s.cellSet, tip)
}

func (s *Snake) addHead(h snakeapp.Cell) {
	s.cellList = append([]snakeapp.Cell{h}, s.cellList...)
	s.cellSet[h] = true
}
