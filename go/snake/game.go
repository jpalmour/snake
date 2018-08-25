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
	snake                     *snake
	food                      cell
}

// New returns a Game with a size by size grid with speed milliseconds per turn.
func New(size, speed int) *Game {
	g := &Game{
		size:  size,
		speed: speed,
		turns: 0,
		score: 0,
	}
	g.populateSnake()
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
		if g.snake.move(getDirection(), g.food) {
			g.populateFood()
			g.score++
		}
		time.Sleep(time.Duration(g.speed) * time.Millisecond)
	}
	fmt.Println()
}

func (g *Game) populateSnake() {
	g.snake = &snake{
		cellList:  []cell{},
		cellSet:   map[cell]bool{},
		direction: up,
	}
	g.snake.addHead(cell{g.size / 2, g.size / 2})
}

func (g *Game) populateFood() {
	food := cell{rand.Intn(g.size), rand.Intn(g.size)}
	for g.snake.bodyCollision(food) {
		food = cell{rand.Intn(g.size), rand.Intn(g.size)}
	}
	g.food = food
}

func (g *Game) finished() bool {
	if g.snake.head().x < 0 || g.snake.head().y < 0 || g.snake.head().x >= g.size || g.snake.head().y >= g.size {
		return true
	}
	if g.snake.headCollision() {
		return true
	}
	return false
}

func getDirection() int {
	// TODO: get direction from keypress
	//consoleReader := bufio.NewReaderSize(os.Stdin, 1)
	//asci, _ := consoleReader.ReadByte()
	//fmt.Println(asci)
	directions := []int{up, down, left, right, none, none, none}
	return directions[rand.Intn(len(directions))]
	// k up 107
	// j down 106
	// l right 108
	// h left 104
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

func clearTerminal() {
	print("\033[H\033[2J")
}
