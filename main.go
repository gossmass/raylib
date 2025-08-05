package main

import (
	"raylib/interp"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	circle_x := interp.New[int32](10)
	circle_y := interp.New[int32](10)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeySpace) {
			pos := rl.GetMousePosition()
			circle_x.SetValue(int32(pos.X)).SetDuration(1.5).SetTransition(interp.InBack)
			circle_y.SetValue(int32(pos.Y)).SetDuration(1.5).SetTransition(interp.OutBack)
		}
		rl.DrawCircle(circle_x.GetValue(), circle_y.GetValue(), 16.0, rl.DarkPurple)

		rl.EndDrawing()
	}
}
