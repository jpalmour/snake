package main

import "github.com/jpalmour/snake/go/snake"

func main() {
	game := snake.New(25, 500)
	game.Play()
}
