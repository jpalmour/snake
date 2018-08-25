package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jpalmour/snake/go/cell"
	"github.com/jpalmour/snake/go/snake"
)

// Game represents the game of snake.
type Game struct {
	size, speed, score, turns int
	snake                     *snake.Snake
	food                      cell.Cell
}

// New returns a Game with a size by size grid with speed milliseconds per turn.
func New(size, speed int) *Game {
	g := &Game{
		size:  size,
		speed: speed,
		turns: 0,
		score: 0,
		snake: snake.New(size),
	}
	g.populateFood()
	return g
}

// Play starts the game.
func (g *Game) Play() {
	for !g.finished() {
		g.turns++
		clearTerminal()
		g.paintScoreboard()
		g.paintGrid()
		if g.snake.Move(getDirection(), g.food) {
			g.populateFood()
			g.score++
		}
		time.Sleep(time.Duration(g.speed) * time.Millisecond)
	}
	fmt.Println()
}

func (g *Game) populateFood() {
	food := cell.Cell{rand.Intn(g.size), rand.Intn(g.size)}
	for g.snake.BodyCollision(food) {
		food = cell.Cell{rand.Intn(g.size), rand.Intn(g.size)}
	}
	g.food = food
}

func (g *Game) finished() bool {
	if g.snake.Head().X < 0 || g.snake.Head().Y < 0 || g.snake.Head().X >= g.size || g.snake.Head().Y >= g.size {
		return true
	}
	if g.snake.HeadCollision() {
		return true
	}
	return false
}

func getDirection() int {
	// TODO: get direction from keypress
	//consoleReader := bufio.NewReaderSize(os.Stdin, 1)
	//asci, _ := consoleReader.ReadByte()
	//fmt.Println(asci)
	directions := []int{cell.Up, cell.Down, cell.Left, cell.Right, cell.None, cell.None, cell.None}
	return directions[rand.Intn(len(directions))]
	// k up 107
	// j down 106
	// l right 108
	// h left 104
}

func (g *Game) paintCell(r, c int) {
	cu := cell.Cell{r, c}
	// TODO: game shouldn't paint snake, snake should (CellSet shouldn't be exported?)
	if g.snake.CellSet[cu] {
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

func clearTerminal() {
	print("\033[H\033[2J")
}
