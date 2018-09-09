package browser

import (
	"fmt"
	"syscall/js"

	"github.com/jpalmour/snake/go"
)

type Controller struct {
	direction int
}

func NewController() *Controller {
	c := &Controller{}
	go c.updateLastPressed()
	return c
}

func (c *Controller) updateLastPressed() {
	document := js.Global().Get("document")
	updateDirection := js.NewEventCallback(js.PreventDefault, func(event js.Value) {
		key := event.Get("key").String()
		fmt.Println("key pressed:", key)
		switch key {
		case "ArrowLeft":
			c.direction = snakeapp.Left
		case "ArrowRight":
			c.direction = snakeapp.Right
		case "ArrowUp":
			c.direction = snakeapp.Up
		case "ArrowDown":
			c.direction = snakeapp.Down
		}
	})
	document.Call("addEventListener", "keydown", updateDirection)
	// TODO: call updateDirection.Release()
}

func (c *Controller) GetDirection() int {
	fmt.Println("direction:", c.direction)
	return c.direction
	//return []int{snakeapp.Up, snakeapp.Down, snakeapp.Left, snakeapp.Right}[rand.Intn(4)]
}
