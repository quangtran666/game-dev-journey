package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Triangle struct {
	A rl.Vector2
	B rl.Vector2
	C rl.Vector2
}

func midpoint(a, b rl.Vector2) rl.Vector2 {
	return rl.NewVector2((a.X+b.X)/2, (a.Y+b.Y)/2)
}

func main() {
	const windowWidth = 800
	const windowHeight = 450

	rl.InitWindow(windowWidth, windowHeight, "game dev journey")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Violet)

		A := rl.NewVector2(400, 100)
		B := rl.NewVector2(200, 300)
		C := rl.NewVector2(600, 300)

		rl.DrawTriangle(A, B, C, rl.Blue)
		rl.DrawCircleV(A, 5, rl.Red)
		rl.DrawCircleV(B, 5, rl.Yellow)
		rl.DrawCircleV(C, 5, rl.Brown)

		centerOfMass := rl.Vector2Scale(rl.Vector2Add(C, rl.Vector2Add(A, B)), 1/3.0)
		rl.DrawLineV(rl.NewVector2(200, 100), centerOfMass, rl.Purple)

		rl.EndDrawing()
	}
}
