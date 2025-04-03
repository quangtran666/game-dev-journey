package character

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Character represents a game character with position and animation
type Character struct {
	Position     rl.Vector2
	Texture      rl.Texture2D
	FrameWidth   int
	FrameHeight  int
	CurrentFrame int
	FrameCount   int
	FrameSpeed   int
	FrameCounter int
	BombRadius   float32
	IsInZone     bool
}

// NewCharacter creates a new character with the given texture
func NewCharacter(texturePath string, frameCount int) Character {
	texture := rl.LoadTexture(texturePath)
	frameWidth := int(texture.Width) / frameCount

	return Character{
		Position:     rl.NewVector2(400, 225), // Center of screen
		Texture:      texture,
		FrameWidth:   frameWidth,
		FrameHeight:  int(texture.Height),
		CurrentFrame: 0,
		FrameCount:   frameCount,
		FrameSpeed:   8, // Adjust for animation speed
		FrameCounter: 0,
		BombRadius:   100,
	}
}

// Update updates the character animation
func (c *Character) Update() {
	c.FrameCounter++
	if c.FrameCounter >= 60/c.FrameSpeed {
		c.FrameCounter = 0
		c.CurrentFrame = (c.CurrentFrame + 1) % c.FrameCount
	}

	speed := float32(4)

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		c.Position.X += speed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		c.Position.X -= speed
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		c.Position.Y -= speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		c.Position.Y += speed
	}

	rl.DrawText(fmt.Sprintf("IsInZone: %v", c.IsInZone), 10, 10, 20, rl.Black)
}

func (c *Character) CheckIsInZone(DummyPosition rl.Vector2) {
	characterToDummy := rl.NewVector2(DummyPosition.X-c.Position.X, DummyPosition.Y-c.Position.Y)
	length := math.Sqrt(float64(characterToDummy.X*characterToDummy.X + characterToDummy.Y*characterToDummy.Y))
	c.IsInZone = length <= float64(c.BombRadius)
}

// Draw draws the character
func (c *Character) Draw() {
	sourceRec := rl.NewRectangle(float32(c.CurrentFrame*c.FrameWidth), 0, float32(c.FrameWidth), float32(c.FrameHeight))
	destRec := rl.NewRectangle(c.Position.X, c.Position.Y, float32(c.FrameWidth), float32(c.FrameHeight))
	origin := rl.NewVector2(float32(c.FrameWidth)/2, float32(c.FrameHeight)/2)

	rl.DrawTexturePro(c.Texture, sourceRec, destRec, origin, 0, rl.White)

	var colorToDraw rl.Color
	if c.IsInZone {
		colorToDraw = rl.Red
	} else {
		colorToDraw = rl.Green
	}

	// Draw bomb radius
	rl.DrawCircleLinesV(c.Position, c.BombRadius, colorToDraw)
}
