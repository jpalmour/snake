package main

import (
	"github.com/jpalmour/snake/go/browser"
	"github.com/jpalmour/snake/go/game"
)

// build with `GOARCH=wasm GOOS=js go build -o snake.wasm main.go`
func main() {
	println("Starting Snake written in Go compiled to WebAssembly...")
	c := browser.NewController()
	d := browser.NewDisplay()
	g := game.New(30, 200, c, d)
	g.Play()
}
