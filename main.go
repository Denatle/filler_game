package main

import (
	"filler/logic"
	"filler/render"
	"filler/shared"
)

func main() {
	var state shared.GameState = shared.GameState{}
	var logic shared.Logic = logic.Logic{}
	var renderer shared.Renderer = render.Renderer{}
	logic.Init(&state)
	renderer.Init(&state)
	for {
		renderer.Render()
	}
}
