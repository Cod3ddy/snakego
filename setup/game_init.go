package setup

import (
	"fmt"

	"github.com/cod3ddy/snakego/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func StartGame() {
	rl.InitWindow(render.ScreenWidth, render.ScreenHeight, "SnakeGo")

	render.InitGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if render.IsStartMenu {
			fmt.Println("We're on the start menu now")
		}

		render.UpdateDrawFrame()
	}
	rl.CloseWindow()
}