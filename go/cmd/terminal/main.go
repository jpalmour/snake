package main

import (
	"github.com/jpalmour/snake/go/game"
	"github.com/jpalmour/snake/go/terminal"
)

func main() {
	c := terminal.NewController()
	d := terminal.NewDisplay()
	g := game.New(35, 100, c, d)
	g.Play()
}
