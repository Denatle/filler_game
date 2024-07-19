package render

import (
	shared "filler/shared"
	"fmt"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	state *shared.GameState
}

func (r Renderer) Render() error {
	raylib.BeginDrawing()

	raylib.ClearBackground(raylib.RayWhite)
	fmt.Println(r.state.Grid[0][0])
	raylib.DrawText("Congrats! You created your first window!", 190, 200, 20, raylib.LightGray)

	raylib.EndDrawing()
	return nil
}

func (r Renderer) Init(state *shared.GameState) {
	r.state = state
	raylib.InitWindow(800, 450, "raylib [core] example - basic window")
	// raylib.SetTargetFPS(60)
}

func (r Renderer) Close() {
	raylib.CloseWindow()
}
