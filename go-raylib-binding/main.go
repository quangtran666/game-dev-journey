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

	maxReflections := 3

	for !rl.WindowShouldClose() {
		character.Update()
		character.CheckWallCollision(walls)

		mousePos := rl.GetMousePosition()
		direction := rl.Vector2Subtract(mousePos, character.Position)
		startingPoint := character.Position

		rl.BeginDrawing()
		rl.ClearBackground(rl.Violet)

		for _, wall := range walls {
			wall.Draw()
		}

		for range maxReflections {
			rayStart, rayEnd := character.CastRay(startingPoint, direction, 1000)
			hit, hitPoint, normal, _ := character.CheckRayWallCollision(rayStart, rayEnd, walls)

			if hit {
				rl.DrawCircleV(hitPoint, 5, rl.Red)
				normalEnd := rl.Vector2Add(hitPoint, rl.Vector2Scale(normal, 100))
				rl.DrawLineV(hitPoint, normalEnd, rl.Blue)
				reflect := rl.Vector2Subtract(direction, rl.Vector2Scale(rl.Vector2Scale(normal, normal.X*direction.X+normal.Y*direction.Y), 2))
				reflectEnd := rl.Vector2Add(hitPoint, rl.Vector2Scale(reflect, 100))
				rl.DrawLineV(hitPoint, reflectEnd, rl.Green)
				direction = reflect
				startingPoint = hitPoint
			}
		}

		character.Draw()
		dummy.Draw()
		character.CheckIsInZone(dummy.Position)

		rl.DrawFPS(25, 25)
		rl.DrawText(fmt.Sprintf("IsInzone: %v", character.IsInZone), 25, 50, 16, rl.Blue)
		rl.EndDrawing()
	}
}
