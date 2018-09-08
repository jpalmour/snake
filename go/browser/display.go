package browser

import (
	"fmt"
	"syscall/js"

	"github.com/jpalmour/snake/go"
)

type Display struct{}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Paint(g *snakeapp.Game) {
	paintScoreboard(g)
	clearGrid(g.Size)
	paintSnake(g.Snake)
	paintFood(g.Food)
}

func clearGrid(s int) {
	document := js.Global().Get("document")
	// TODO: Does this create a memory leak? Learn what remove does with removed resources.
	document.Call("getElementById", "grid").Call("remove")
	game := document.Call("getElementById", "game")
	grid := document.Call("createElement", "div")
	grid.Set("id", "grid")
	game.Call("appendChild", grid)
	// TODO: set template columns/rows to s instead of relying on static css values
}

func paintScoreboard(g *snakeapp.Game) {
	// TODO: add scoreboard to HTML
	fmt.Printf("Snake\t\tScore: %d\t\tSpeed: %d\t\tTurns: %d\n", g.Score, g.Speed, g.Turns)
}

func paintFood(c snakeapp.Cell) {
	document := js.Global().Get("document")
	grid := document.Call("getElementById", "grid")
	food := document.Call("createElement", "div")
	food.Get("style").Set("grid-column", js.ValueOf(c.X+1))
	food.Get("style").Set("grid-row", js.ValueOf(c.Y+1))
	food.Get("style").Set("background-color", js.ValueOf("red"))
	grid.Call("appendChild", food)

}

func paintSnake(s snakeapp.Snake) {
	document := js.Global().Get("document")
	for _, c := range s.CellList() {
		grid := document.Call("getElementById", "grid")
		snakeCell := document.Call("createElement", "div")
		snakeCell.Get("style").Set("grid-column", js.ValueOf(c.X+1))
		snakeCell.Get("style").Set("grid-row", js.ValueOf(c.Y+1))
		snakeCell.Get("style").Set("background-color", js.ValueOf("black"))
		grid.Call("appendChild", snakeCell)
	}
}
