package snake

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Game represents the game of snake.
type Game struct {
	size, speed, score, turns int
	snake                     snake
	food                      *cell
}

// snake represents the snake that the player controls.
type snake []*cell

// cell represents a location in Game's grid.
type cell struct {
	x, y int
}

// New returns a Game with a size by size grid with speed milliseconds per turn.
func New(size, speed int) *Game {
	return &Game{
		size:  size,
		speed: speed,
		snake: []*cell{},
		turns: 0,
		score: 0,
	}
}

// Play starts the game.
func (g *Game) Play() {
	g.populateSnake()
	g.generateFood()
	for !g.finished() {
		g.turns++
		clearTerminal()
		g.paintScoreboard()
		g.paintGrid()
		time.Sleep(time.Duration(g.speed) * time.Millisecond)
	}
	fmt.Println()
}

func (g *Game) finished() bool {
	return false
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

func (g *Game) paintCell(r, c int) {
	fmt.Print("  ")
}

func (g *Game) paintBorder() {
	fmt.Printf("*%s*\n", strings.Repeat("-", 2*g.size))
}

func (g *Game) populateSnake() {
	g.snake = []*cell{&cell{rand.Intn(g.size), rand.Intn(g.size)}}
}

func (g *Game) generateFood() {
	if g.food == nil {
		// TODO: ensure food not on snake
		g.food = &cell{rand.Intn(g.size), rand.Intn(g.size)}
	}
}

func clearTerminal() {
	print("\033[H\033[2J")
}
