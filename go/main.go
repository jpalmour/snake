package main

import "github.com/jpalmour/snake/go/snake"

func main() {
	game := snake.New(10, 500)
	game.Play()
}
