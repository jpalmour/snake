package snake

import (
	"math/rand"

	"github.com/jpalmour/snake/go/cell"
)

// Snake represents the snake that the player controls.
type Snake struct {
	cellList  []cell.Cell
	CellSet   map[cell.Cell]bool
	direction int
}

func New(boardSize int) *Snake {
	s := &Snake{
		cellList:  []cell.Cell{},
		CellSet:   map[cell.Cell]bool{},
		direction: cell.Up,
	}
	// TODO: should addHead be exported?
	s.addHead(cell.Cell{boardSize / 2, boardSize / 2})
	return s
}

func (s *Snake) BodyCollision(c cell.Cell) bool {
	return contains(s.cellList[0:], c)
}

func (s *Snake) HeadCollision() bool {
	return contains(s.cellList[1:], s.Head())
}

func contains(l []cell.Cell, c cell.Cell) bool {
	for _, ci := range l {
		if c == ci {
			return true
		}
	}
	return false
}

func (s *Snake) Head() cell.Cell {
	return s.cellList[0]
}

func opposite(d1, d2 int) bool {
	type pair struct {
		a, b int
	}
	ds := pair{d1, d2}
	return ds == pair{cell.Up, cell.Down} || ds == pair{cell.Down, cell.Up} || ds == pair{cell.Left, cell.Right} || ds == pair{cell.Right, cell.Left}
}

func (s *Snake) getDirection(d int) int {
	if d == cell.None || opposite(d, s.direction) {
		return s.direction
	}
	return d
}

func (s *Snake) getNewHead(d int) cell.Cell {
	switch d {
	case cell.Down:
		return cell.Cell{s.Head().X, s.Head().Y + 1}
	case cell.Left:
		return cell.Cell{s.Head().X - 1, s.Head().Y}
	case cell.Right:
		return cell.Cell{s.Head().X + 1, s.Head().Y}
	}
	return cell.Cell{s.Head().X, s.Head().Y - 1}
}

func (s *Snake) Move(d int, food cell.Cell) bool {
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

func (s *Snake) removeTailTip() {
	tip := s.cellList[len(s.cellList)-1]
	s.cellList = s.cellList[0 : len(s.cellList)-1]
	delete(s.CellSet, tip)
}

func (s *Snake) addHead(h cell.Cell) {
	s.cellList = append([]cell.Cell{h}, s.cellList...)
	s.CellSet[h] = true
}
