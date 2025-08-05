package main

import (
	"raylib/interp"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getRaio() float64 {
	screenWidth := float64(rl.GetScreenWidth())
	screenHeight := float64(rl.GetScreenHeight())
	return screenWidth / screenHeight
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	circle_x := interp.New[int32](10)
	circle_y := interp.New[int32](10)

	camera := rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)

		if rl.IsWindowResized() {
			camera.Zoom = float32(getRaio())
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)

			circle_x.SetValue(int32(pos.X)).SetDuration(1.5).SetTransition(interp.OutBack)
			circle_y.SetValue(int32(pos.Y)).SetDuration(1.5).SetTransition(interp.OutBack)
		}
		rl.DrawCircle(circle_x.GetValue(), circle_y.GetValue(), 16.0, rl.DarkPurple)
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
