package terminal

import (
	"fmt"
	"strings"

	"github.com/jpalmour/snake/go"
)

type Display struct {
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Paint(g *snakeapp.Game) {
	clearTerminal()
	d.paintScoreboard(g)
	d.paintGrid(g)
}

func (d *Display) paintScoreboard(g *snakeapp.Game) {
	fmt.Printf("Snake (written in Go)\t\tScore: %d\t\tSpeed: %d\t\tTurns: %d\n", g.Score, g.Speed, g.Turns)
}

func (d *Display) paintGrid(g *snakeapp.Game) {
	d.paintBorder(g)
	for r := 0; r < g.Size; r++ {
		d.paintRow(r, g)
	}
	d.paintBorder(g)
}

func (d *Display) paintCell(r, c int, g *snakeapp.Game) {
	cu := snakeapp.Cell{r, c}
	if g.Snake.Cells()[cu] {
		fmt.Print("@")
	} else if cu == g.Food {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func (d *Display) paintRow(r int, g *snakeapp.Game) {
	fmt.Print("|")
	for c := 0; c < g.Size; c++ {
		d.paintCell(r, c, g)
	}
	fmt.Println("|")
}

func (d *Display) paintBorder(g *snakeapp.Game) {
	fmt.Printf("*%s*\n", strings.Repeat("-", g.Size))
}

func clearTerminal() {
	print("\033[H\033[2J")
}
