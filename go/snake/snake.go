package snake

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	up = iota
	down
	left
	right
)

// Game represents the game of snake.
type Game struct {
	size, speed, score, turns, direction int
	snake                                *snake
	food                                 cell
}

// snake represents the snake that the player controls.
type snake struct {
	cellList []cell
	cellSet  map[cell]bool
}

func (g *Game) populateSnake() {
	g.snake = &snake{
		cellList: []cell{},
		cellSet:  map[cell]bool{},
	}
	g.snake.addHead(cell{g.size / 2, g.size / 2})
}

// cell represents a location in Game's grid.
type cell struct {
	x, y int
}

// New returns a Game with a size by size grid with speed milliseconds per turn.
func New(size, speed int) *Game {
	g := &Game{
		size:      size,
		speed:     speed,
		turns:     0,
		score:     0,
		direction: right,
	}
	g.populateSnake()
	g.generateFood()
	return g
}

// Play starts the game.
func (g *Game) Play() {
	for !g.finished() {
		g.turns++
		clearTerminal()
		g.paintScoreboard()
		g.paintGrid()
		g.updateSnake()
		time.Sleep(time.Duration(g.speed) * time.Millisecond)
	}
	fmt.Println()
}

func (g *Game) finished() bool {
	return false
}

func (g *Game) updateSnake() {
	g.direction = getDirection()
	// TODO: fix law of demeter violation
	if g.snake.move(g.direction, g.food) {
		g.score++
	}

}

func (s *snake) head() cell {
	return s.cellList[0]
}

func (s *snake) move(d int, foodLoc cell) bool {
	newHead := cell{s.head().x - 1, s.head().y}
	if d == up {
		newHead = cell{s.head().x, s.head().y - 1}
	} else if d == down {
		newHead = cell{s.head().x, s.head().y + 1}
	} else if d == right {
		newHead = cell{s.head().x + 1, s.head().y}
	}
	eatsFood := foodLoc == newHead
	if !eatsFood {
		s.removeTailTip()
	}
	s.addHead(newHead)
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

func getDirection() int {
	// TODO: get direction from keypress
	return right
}

func (g *Game) paintCell(r, c int) {
	cu := cell{r, c}
	if g.snake.cellSet[cu] {
		fmt.Print("@")
	} else if cu == g.food {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func (g *Game) paintScoreboard() {
	fmt.Printf("Snake (written in Go)\t\tScore: %d\t\tSpeed: %d\t\tTurns: %d\n", g.score, g.speed, g.turns)
}

func (g *Game) paintGrid() {
	g.paintBorder()
	for r := 0; r < g.size; r++ {
		g.paintRow(r)
	}
	g.paintBorder()
}

func (g *Game) paintRow(r int) {
	fmt.Print("|")
	for c := 0; c < g.size; c++ {
		g.paintCell(r, c)
	}
	fmt.Println("|")
}

func (g *Game) paintBorder() {
	fmt.Printf("*%s*\n", strings.Repeat("-", g.size))
}

func (g *Game) generateFood() {
	// TODO: don't use a valid cell for detecting missing food
	c := cell{0, 0}
	if g.food == c {
		// TODO: ensure food not on snake
		g.food = cell{rand.Intn(g.size), rand.Intn(g.size)}
	}
}

func clearTerminal() {
	print("\033[H\033[2J")
}
