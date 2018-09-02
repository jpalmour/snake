package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jpalmour/snake/go"
	"github.com/jpalmour/snake/go/snake"
)

// New returns a Game with a Size by Size grid with Speed milliseconds per turn.
func New(size, speed int, c snakeapp.Controller, d snakeapp.Display) *Game {
	g := &Game{
		Size:  size,
		Speed: speed,
		Turns: 0,
		Score: 0,
		Snake: snake.New(size, c),
	}

	g.populateFood()
	return g
}

// Play starts the game.
func (g *Game) Play() {
	for !g.finished() {
		g.Turns++
		g.Display.Paint(&g)
		if g.Snake.Move(g.Food) {
			g.populateFood()
			g.Score++
		}
		time.Sleep(time.Duration(g.Speed) * time.Millisecond)
	}
	fmt.Println()
}

func (g *Game) populateFood() {
	food := snakeapp.Cell{rand.Intn(g.Size), rand.Intn(g.Size)}
	for g.Snake.BodyCollision(food) {
		food = snakeapp.Cell{rand.Intn(g.Size), rand.Intn(g.Size)}
	}
	g.Food = food
}

func (g *Game) finished() bool {
	if g.Snake.Head().X < 0 || g.Snake.Head().Y < 0 || g.Snake.Head().X >= g.Size || g.Snake.Head().Y >= g.Size {
		return true
	}
	if g.Snake.HeadCollision() {
		return true
	}
	return false
}
