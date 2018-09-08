package snakeapp

// Game represents the game of snake.
type Game struct {
	Size, Speed, Score, Turns int
	Snake                     Snake
	Food                      Cell
	Display                   Display
}

// Cell represents a location in Game's grid.
type Cell struct {
	X, Y int
}

type Controller interface {
	GetDirection() int
}

type Display interface {
	Paint(*Game)
}

type Snake interface {
	BodyCollision(Cell) bool
	HeadCollision() bool
	Head() Cell
	Move(Cell) bool
	// TODO: remove CellSet
	CellSet() map[Cell]bool
	CellList() []Cell
}

const (
	Up = iota
	Down
	Left
	Right
)
