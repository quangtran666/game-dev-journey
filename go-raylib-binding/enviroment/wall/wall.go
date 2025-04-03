package wall

import rl "github.com/gen2brain/raylib-go/raylib"

type Wall struct {
	Rect  rl.Rectangle
	Color rl.Color
}

func NewWall(x, y, width, height float32, color rl.Color) Wall {
	return Wall{
		Rect:  rl.Rectangle{X: x, Y: y, Width: width, Height: height},
		Color: color,
	}
}

func (w *Wall) Draw() {
	rl.DrawRectangleRec(w.Rect, w.Color)
}

func (w *Wall) CheckCollision(characterPosition *rl.Vector2, characterSize rl.Vector2) bool {
	// Tạo hình chữ nhật đại diện cho nhân vật
	characterRect := rl.Rectangle{
		X:      characterPosition.X - characterSize.X/2,
		Y:      characterPosition.Y - characterSize.Y/2,
		Width:  characterSize.X,
		Height: characterSize.Y,
	}

	return rl.CheckCollisionRecs(w.Rect, characterRect)
}
