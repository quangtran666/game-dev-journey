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

	colors := []rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Purple, rl.Orange, rl.Pink}
	triangles := []Triangle{
		{A: rl.NewVector2(400, 100), B: rl.NewVector2(200, 300), C: rl.NewVector2(600, 300)},
	}

	iteration := 6
	for i := 0; i < iteration; i++ {
		numTriangles := len(triangles)
		for j := 0; j < numTriangles; j++ {
			firstTriangle := triangles[0]
			triangles = triangles[1:]

			mAB := midpoint(firstTriangle.A, firstTriangle.B)
			mBC := midpoint(firstTriangle.B, firstTriangle.C)
			mAC := midpoint(firstTriangle.A, firstTriangle.C)

			triangles = append(triangles, Triangle{A: firstTriangle.A, B: mAB, C: mAC})
			triangles = append(triangles, Triangle{A: mAB, B: firstTriangle.B, C: mBC})
			triangles = append(triangles, Triangle{A: mAC, B: mBC, C: firstTriangle.C})
		}
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Violet)

		// Draw all triangles
		for index, triangle := range triangles {
			rl.DrawTriangle(triangle.A, triangle.B, triangle.C, colors[index%len(colors)])
		}

		rl.EndDrawing()
	}
}
