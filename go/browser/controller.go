package browser

import (
	"math/rand"

	"github.com/jpalmour/snake/go"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetDirection() int {
	return []int{snakeapp.Up, snakeapp.Down, snakeapp.Left, snakeapp.Right}[rand.Intn(4)]
}
