package terminal

import (
	"fmt"
	"strings"

	"github.com/jpalmour/snake/go/game"
)

type Display struct {
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Paint(g *game.Game) {
	clearTerminal()
	d.paintScoreboard(g)
	d.paintGrid(g)
}

func (d *Display) paintScoreboard(g *game.Game) {
	fmt.Printf("Snake (written in Go)\t\tScore: %d\t\tSpeed: %d\t\tTurns: %d\n", g.score, g.speed, g.turns)
}

func (d *Display) paintGrid(g *game.Game) {
	d.paintBorder()
	for r := 0; r < g.size; r++ {
		d.paintRow(r)
	}
	d.paintBorder()
}

func (d *Display) paintCell(r, c int) {
	cu := snakeapp.Cell{r, c}
	if g.snake.CellSet[cu] {
		fmt.Print("@")
	} else if cu == g.food {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func (d *Display) paintRow(r int) {
	fmt.Print("|")
	for c := 0; c < g.size; c++ {
		g.paintCell(r, c)
	}
	fmt.Println("|")
}

func (d *Display) paintBorder() {
	fmt.Printf("*%s*\n", strings.Repeat("-", g.size))
}

func clearTerminal() {
	print("\033[H\033[2J")
}
