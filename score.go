package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func displayScore(score int) {
	rl.ClearBackground(rl.RayWhite)
	scoreText := fmt.Sprintf("Score: %d", score)
	var fontSize int32 = 50
	rl.DrawText(
		scoreText,
		int32(rl.GetScreenWidth())-int32(float64(99/100.0)*float64(rl.GetScreenWidth())),   // posX
		int32(rl.GetScreenHeight())-int32(float64(99/100.0)*float64(rl.GetScreenHeight())), // posY
		fontSize,
		rl.Black,
	)
}
