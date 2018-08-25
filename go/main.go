package main

import "github.com/jpalmour/snake/go/snake"

func main() {
	game := snake.New(35, 500)
	game.Play()
}
