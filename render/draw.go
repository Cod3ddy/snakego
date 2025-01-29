package render

import rl "github.com/gen2brain/raylib-go/raylib"

type (
	Drawable interface {
		Draw()
		ZIndex() int
	}

	Text struct {
		Text     string
		FontSize int32
		X, Y     int32
		Color    rl.Color
		Zindex   int
	}

	SnakeSegment struct {
		Position rl.Vector2
		Size     rl.Vector2
		Speed    rl.Vector2
		Color    rl.Color
		Zindex   int
	}

	Food struct {
		Position rl.Vector2
		Size     rl.Vector2
		IsActive bool
		Color    rl.Color
		Zindex   int
	}
)

func (s *SnakeSegment) Draw() {
	rl.DrawRectangleV(s.Position, s.Size, s.Color)
}

func (s *SnakeSegment) ZIndex() int {
	return s.Zindex
}

func (t *Text) Draw() {
	rl.DrawText(t.Text, t.X, t.Y, t.FontSize, t.Color)
}

func (t *Text) ZIndex() int {
	return t.Zindex
}
