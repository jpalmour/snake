package main

import (
	"github.com/jpalmour/snake/go/game"
	"github.com/jpalmour/snake/go/keyboard"
	"github.com/jpalmour/snake/go/terminal"
)

func main() {
	c := keyboard.NewController()
	d := terminal.NewDisplay()
	g := game.New(35, 250, c, d)
	g.Play()
}
