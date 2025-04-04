package character

import (
	"math"

	"mygame/enviroment/wall"

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
}

func (c *Character) CastRay(direction rl.Vector2, maxLength float32) (rl.Vector2, rl.Vector2) {
	normarlize := rl.Vector2Normalize(direction)

	rayEnd := rl.NewVector2(
		c.Position.X+normarlize.X*maxLength,
		c.Position.Y+normarlize.Y*maxLength)

	return c.Position, rayEnd
}

func (c *Character) CheckRayWallCollision(rayStart, rayEnd rl.Vector2, walls []wall.Wall) (bool, rl.Vector2, rl.Vector2, int) {
	closestHit := false
	closestPoint := rl.NewVector2(0, 0)
	cloestDistance := float32(math.MaxFloat32)
	cloestWallIndex := -1

	rl.DrawLineV(rayStart, rayEnd, rl.Brown)
	for i, wall := range walls {
		topLeft := rl.NewVector2(wall.Rect.X, wall.Rect.Y)
		topRight := rl.NewVector2(wall.Rect.X+wall.Rect.Width, wall.Rect.Y)
		bottomLeft := rl.NewVector2(wall.Rect.X, wall.Rect.Y+wall.Rect.Height)
		bottomRight := rl.NewVector2(wall.Rect.X+wall.Rect.Width, wall.Rect.Y+wall.Rect.Height)

		edges := []struct{ start, end rl.Vector2 }{
			{topLeft, topRight},
			{topRight, bottomRight},
			{bottomRight, bottomLeft},
			{bottomLeft, topLeft},
		}

		for _, edge := range edges {
			var collisionPoint rl.Vector2
			if rl.CheckCollisionLines(rayStart, rayEnd, edge.start, edge.end, &collisionPoint) {
				// Tính khoảng cách từ điểm bắt đầu ray đến điểm va chạm
				distance := rl.Vector2Distance(collisionPoint, rayStart)

				if distance < cloestDistance {
					cloestDistance = distance
					closestPoint = collisionPoint
					closestHit = true
					cloestWallIndex = i
				}
			}
		}
	}

	normal := rl.NewVector2(0, 0)
	if closestHit {
		// Dựa vào tường va chạm để xác định pháp tuyến
		wall := walls[cloestWallIndex]

		if math.Abs(float64(closestPoint.X-wall.Rect.X)) < 0.1 {
			normal = rl.NewVector2(-1, 0) // Va chạm với tường trái
		} else if math.Abs(float64(closestPoint.X-(wall.Rect.X+wall.Rect.Width))) < 0.1 {
			normal = rl.NewVector2(1, 0) // Va chạm với tường phải
		} else if math.Abs(float64(closestPoint.Y-wall.Rect.Y)) < 0.1 {
			normal = rl.NewVector2(0, -1) // Va chạm với tường dưới
		} else if math.Abs(float64(closestPoint.Y-(wall.Rect.Y+wall.Rect.Height))) < 0.1 {
			normal = rl.NewVector2(0, 1) // Va chạm với tường trên
		}
	}

	return closestHit, closestPoint, normal, cloestWallIndex
}

func (c *Character) CheckWallCollision(walls []wall.Wall) {
	// Store the old position for reversing if needed
	oldPosition := c.Position

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

	characterSize := c.GetSize()
	for _, w := range walls {
		if w.CheckCollision(&c.Position, characterSize) {
			c.Position = oldPosition // Revert to the old position if a collision occurs
			break                    // Exit the loop after the first collision
		}
	}
}

func (c *Character) GetSize() rl.Vector2 {
	return rl.NewVector2(float32(c.FrameWidth), float32(c.FrameHeight))
}

func (c *Character) CheckIsInZone(DummyPosition rl.Vector2) {
	characterToDummy := rl.NewVector2(DummyPosition.X-c.Position.X, DummyPosition.Y-c.Position.Y)
	length := math.Sqrt(float64(characterToDummy.X*characterToDummy.X + characterToDummy.Y*characterToDummy.Y))
	c.IsInZone = length <= float64(c.BombRadius)
}

// Draw draws the character
func (c *Character) Draw() {
	// Draw the character texture with animation
	sourceRec := rl.NewRectangle(float32(c.CurrentFrame*c.FrameWidth), 0, float32(c.FrameWidth), float32(c.FrameHeight))
	destRec := rl.NewRectangle(c.Position.X, c.Position.Y, float32(c.FrameWidth), float32(c.FrameHeight))
	origin := rl.NewVector2(float32(c.FrameWidth)/2, float32(c.FrameHeight)/2)

	rl.DrawTexturePro(c.Texture, sourceRec, destRec, origin, 0, rl.White)

	// Draw bomb radius
	var colorToDraw rl.Color
	if c.IsInZone {
		colorToDraw = rl.Red
	} else {
		colorToDraw = rl.Green
	}

	rl.DrawCircleLinesV(c.Position, c.BombRadius, colorToDraw)

	// Draw character rectangle for debugging
	rl.DrawRectangleLines(
		int32(c.Position.X-float32(c.FrameWidth/2)),
		int32(c.Position.Y-float32(c.FrameHeight/2)),
		int32(c.FrameWidth),
		int32(c.FrameHeight),
		rl.Blue)
}
