package main

import (
	"raylib/interp"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const EASING_FUNCTIONS = "Linear;InQuad;OutQuad;InOutQuad;InCubic;OutCubic;InOutCubic;InQuart;OutQuart;InOutQuart;InQuint;OutQuint;InOutQuint;InSine;OutSine;InOutSine;InExpo;OutExpo;InOutExpo;InCirc;OutCirc;InOutCirc;InElastic;OutElastic;InOutElastic;InElasticFunction;OutElasticFunction;InOutElasticFunction;InBack;OutBack;InOutBack;InBounce;OutBounce;InOutBounce;InSquare;OutSquare;InOutSquare"

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
	circle_r := interp.New[float32](16.0)
	radiusSwitch := false

	camera := rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 1)
	var dropDownActive int32 = 0
	dropDownEditMode := false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.RayWhite)

		if rl.IsWindowResized() {
			camera.Zoom = float32(getRaio())
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)

			circle_x.SetValue(int32(pos.X)).SetDuration(1.5)
			circle_y.SetValue(int32(pos.Y)).SetDuration(1.5)

			if radiusSwitch {
				circle_r.SetValue(16)
			} else {
				circle_r.SetValue(32)
			}

			circle_r.SetDuration(1.5)
			radiusSwitch = !radiusSwitch
		}
		rl.DrawCircle(circle_x.GetValue(), circle_y.GetValue(), circle_r.GetValue(), rl.DarkPurple)
		rl.EndMode2D()

		if gui.DropdownBox(rl.Rectangle{0, 0, 128, 24}, EASING_FUNCTIONS, &dropDownActive, dropDownEditMode) {
			dropDownEditMode = !dropDownEditMode

			fn := interp.GetEasingFunction(interp.Easing(dropDownActive))
			circle_x.SetTransition(fn)
			circle_y.SetTransition(fn)
			circle_r.SetTransition(fn)
		}

		rl.EndDrawing()
	}
}
