package dummy

import rl "github.com/gen2brain/raylib-go/raylib"

type Dummy struct {
	Position rl.Vector2
	Texture  rl.Texture2D
}

func NewDummy(texturePath string) Dummy {
	return Dummy{
		Position: rl.NewVector2(600, 150),
		Texture:  rl.LoadTexture(texturePath),
	}
}

func (d *Dummy) Draw() {
	sourceRec := rl.NewRectangle(0, 0, float32(d.Texture.Width), float32(d.Texture.Height))
	destRec := rl.NewRectangle(d.Position.X, d.Position.Y, float32(d.Texture.Width), float32(d.Texture.Height))
	origin := rl.NewVector2(float32(d.Texture.Width)/2, float32(d.Texture.Height)/2)

	rl.DrawTexturePro(d.Texture, sourceRec, destRec, origin, 0, rl.White)
}
