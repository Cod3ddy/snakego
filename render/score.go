package render

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameScore(score int) *GameText{
	scoreText := fmt.Sprintf("Score: %d", score)
	textX := int32(rl.GetScreenWidth())-int32(float64(99/100.0)*float64(rl.GetScreenWidth()))
	textY :=  int32(rl.GetScreenHeight())-int32(float64(99/100.0)*float64(rl.GetScreenHeight()))
	var fontSize int32 = 20
	var zIndex int = 10

	return &GameText{
		Text: scoreText,
		X: textX, 
		Y: textY, 
		FontSize: 
		fontSize, 
		Zindex: zIndex, 	
		Color: rl.Black,
	}
}
