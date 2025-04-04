package main

import (
	"fmt"

	"mygame/entity/character"
	"mygame/entity/dummy"
	"mygame/enviroment/wall"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const windowWidth = 800
	const windowHeight = 450

	rl.InitWindow(windowWidth, windowHeight, "game dev journey")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	character := character.NewCharacter("C:/Projects/game-dev-journey/go-raylib-binding/assets/bomb_character_o_idle.png", 2)
	defer rl.UnloadTexture(character.Texture)
	dummy := dummy.NewDummy("C:/Projects/game-dev-journey/go-raylib-binding/assets/Dummy.png")
	defer rl.UnloadTexture(dummy.Texture)

	wallThickness := float32(20)
	walls := []wall.Wall{
		wall.NewWall(0, 0, windowWidth, wallThickness, rl.DarkGray),
		wall.NewWall(0, windowHeight-wallThickness, windowWidth, wallThickness, rl.DarkGray),
		wall.NewWall(0, 0, wallThickness, windowHeight, rl.DarkGray),
		wall.NewWall(windowWidth-wallThickness, 0, wallThickness, windowHeight, rl.DarkGray),
	}

	for !rl.WindowShouldClose() {
		character.Update()
		character.CheckWallCollision(walls)

		mousePos := rl.GetMousePosition()
		direction := rl.Vector2Subtract(mousePos, character.Position)

		rayStart, rayEnd := character.CastRay(direction, 1000)
		hit, hitPoint, normal, wallIndex := character.CheckRayWallCollision(rayStart, rayEnd, walls)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Violet)

		for _, wall := range walls {
			wall.Draw()
		}

		if hit {
			// Draw hit point
			rl.DrawCircleV(hitPoint, 5, rl.Red)
			// Draw normal line
			normalEnd := rl.Vector2Add(hitPoint, rl.Vector2Scale(normal, 200))
			rl.DrawLineV(hitPoint, normalEnd, rl.Blue)

			rl.DrawText(fmt.Sprintf("Hit Wall: %d", wallIndex), 25, 75, 16, rl.Blue)
			rl.DrawText(fmt.Sprintf("Normal: x: %g y: %g", normal.X, normal.Y), 25, 100, 16, rl.Blue)
		}

		character.Draw()
		dummy.Draw()
		character.CheckIsInZone(dummy.Position)

		rl.DrawFPS(25, 25)
		rl.DrawText(fmt.Sprintf("IsInzone: %v", character.IsInZone), 25, 50, 16, rl.Blue)
		rl.EndDrawing()
	}
}
