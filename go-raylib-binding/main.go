package main

import (
	"mygame/entity/character"
	"mygame/entity/dummy"

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

	for !rl.WindowShouldClose() {
		character.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Violet)

		character.Draw()
		dummy.Draw()
		character.CheckIsInZone(dummy.Position)

		rl.EndDrawing()
	}
}
