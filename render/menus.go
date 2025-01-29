package render

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameScoreUI(score int) *GameText{
	scoreString := fmt.Sprintf("Score: %d", score)
	textX := int32(rl.GetScreenWidth())-int32(float64(99/100.0)*float64(rl.GetScreenWidth()))
	textY :=  int32(rl.GetScreenHeight())-int32(float64(99/100.0)*float64(rl.GetScreenHeight()))
	var fontSize int32 = 20
	var zIndex int = 10

	return &GameText{
		Text: scoreString,
		X: textX, 
		Y: textY, 
		FontSize: 
		fontSize, 
		Zindex: zIndex, 	
		Color: rl.Black,
	}
}

func GamePausedUI(screenWidth, screenHeight int32) *GameText{
	pauseString := "Game Paused"
	var fontSize int32 = 40
	textX :=  screenWidth/2-rl.MeasureText(pauseString, fontSize)/2
	textY :=  screenHeight/2-fontSize
	fontColor := rl.Gray
	zIndex := 2

	return &GameText{
		Text: pauseString, 
		X: textX,
		Y: textY, 
		Color: fontColor, 
		Zindex: zIndex,
		FontSize: fontSize,
	}
}
