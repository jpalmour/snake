package terminal

import (
	"github.com/jpalmour/snake/go"
	termbox "github.com/nsf/termbox-go"
)

type Controller struct {
	direction int
}

func NewController() *Controller {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	c := &Controller{
		direction: snakeapp.Right,
	}
	// TODO: Share memory by communicating; don't communicate by sharing memory.
	go c.updateLastPressed()
	return c
}

func (c *Controller) GetDirection() int {
	return c.direction
}

func (c *Controller) updateLastPressed() {
	defer termbox.Close()
	for {
		ev := termbox.PollEvent()
		switch ev.Key {
		case termbox.KeyArrowLeft:
			c.direction = snakeapp.Left
		case termbox.KeyArrowRight:
			c.direction = snakeapp.Right
		case termbox.KeyArrowUp:
			c.direction = snakeapp.Up
		case termbox.KeyArrowDown:
			c.direction = snakeapp.Down
		}
	}
}
