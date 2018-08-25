package main

import "github.com/jpalmour/snake/go/game"

func main() {
	g := game.New(35, 500)
	g.Play()
}
