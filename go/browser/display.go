package browser

import (
	"fmt"
	"syscall/js"

	"github.com/jpalmour/snake/go"
)

type Display struct {
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Paint(g *snakeapp.Game) {
	d.paintScoreboard(g)
	d.paintGrid(g)
}

func (d *Display) paintScoreboard(g *snakeapp.Game) {
	fmt.Printf("Snake\t\tScore: %d\t\tSpeed: %d\t\tTurns: %d\n", g.Score, g.Speed, g.Turns)
}

func (d *Display) paintGrid(g *snakeapp.Game) {
	for r := 0; r < g.Size; r++ {
		d.paintRow(r, g)
	}
}

func (d *Display) paintRow(r int, g *snakeapp.Game) {
	for c := 0; c < g.Size; c++ {
		d.paintCell(c, r, g)
	}
}

func (d *Display) paintCell(r, c int, g *snakeapp.Game) {
	cu := snakeapp.Cell{r, c}
	if g.Snake.Cells()[cu] {
		document := js.Global().Get("document")
		grid := document.Call("getElementById", "grid")
		snakeCell := document.Call("createElement", "div")
		snakeCell.Get("style").Set("grid-column", js.ValueOf(c+1))
		snakeCell.Get("style").Set("grid-row", js.ValueOf(r+1))
		snakeCell.Get("style").Set("background-color", js.ValueOf("black"))
		grid.Call("appendChild", snakeCell)
	} else if cu == g.Food {
		document := js.Global().Get("document")
		grid := document.Call("getElementById", "grid")
		snakeCell := document.Call("createElement", "div")
		snakeCell.Get("style").Set("grid-column", js.ValueOf(c+1))
		snakeCell.Get("style").Set("grid-row", js.ValueOf(r+1))
		snakeCell.Get("style").Set("background-color", js.ValueOf("red"))
		grid.Call("appendChild", snakeCell)
	} else {
		//fmt.Print(" ")
	}
}
