package snakeapp

import (
	"github.com/jpalmour/snake/go/snake"
)

// Game represents the game of snake.
type Game struct {
	Size, Speed, Score, Turns int
	Snake                     *snake.Snake
	Food                      snakeapp.Cell
	Display                   snakeapp.Display
}

// Cell represents a location in Game's grid.
type Cell struct {
	X, Y int
}

type Controller interface {
	GetDirection() int
}

type Display interface {
	Paint(*snakeapp.Game)
}

const (
	Up = iota
	Down
	Left
	Right
	None
)
