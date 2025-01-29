package main

import (
	"fmt"

	"github.com/cod3ddy/snakego/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	snakeLength = 256
	squareSize  = 31

	screenWidth  = 800
	screenHeight = 500
)

var (
	gameOver     bool = false
	frameCounter int  = 0
	pause        bool = false
	allowMove    bool = false
	isStartMenu  bool = false
	score        int  = 0
	counterTail  int  = 0
	offset       rl.Vector2
)

type (
	SnakeSegment struct {
		Position rl.Vector2
		Size     rl.Vector2
		Speed    rl.Vector2
		Color    rl.Color
	}

	Food struct {
		Position rl.Vector2
		Size     rl.Vector2
		IsActive bool
		Color    rl.Color
	}
)

var (
	fruit         render.Food
	snake         []render.SnakeSegment = make([]render.SnakeSegment, snakeLength)
	snakePosition []rl.Vector2   = make([]rl.Vector2, snakeLength)
)

func InitGame() {
	frameCounter = 0
	gameOver = false
	pause = false
	isStartMenu = false

	counterTail = 1
	allowMove = false

	offset.X = screenWidth % squareSize
	offset.Y = screenHeight % squareSize

	for i := 0; i < snakeLength; i++ {
		snake[i].Position = rl.Vector2{X: offset.X / 2, Y: offset.Y / 2}
		snake[i].Size = rl.Vector2{X: squareSize, Y: squareSize}
		snake[i].Speed = rl.Vector2{X: squareSize, Y: 0}

		if i == 0 {
			snake[i].Color = rl.DarkBlue
		} else {
			snake[i].Color = rl.Blue
		}
	}

	for i := 0; i < snakeLength; i++ {
		snakePosition[i] = rl.Vector2{X: 0.0, Y: 0.0}
	}

	fruit.Size = rl.Vector2{X: squareSize, Y: squareSize}
	fruit.Color = rl.SkyBlue
	fruit.IsActive = false
	displayScore(score)
}

func UpdateGame() {
	if !gameOver {
		if rl.IsKeyPressed('P') {
			pause = !pause
		}

		if rl.IsKeyPressed(rl.KeyM) {
			isStartMenu = !isStartMenu
		}

		if !pause && !isStartMenu {
			// Player controls
			if rl.IsKeyPressed(rl.KeyRight) && snake[0].Speed.X == 0 && allowMove {
				snake[0].Speed = rl.Vector2{X: squareSize, Y: 0}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyLeft) && snake[0].Speed.X == 0 && allowMove {
				snake[0].Speed = rl.Vector2{X: -squareSize, Y: 0}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyUp) && snake[0].Speed.Y == 0 && allowMove {
				snake[0].Speed = rl.Vector2{X: 0, Y: -squareSize}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyDown) && snake[0].Speed.Y == 0 && allowMove {
				snake[0].Speed = rl.Vector2{X: 0, Y: squareSize}
				allowMove = false
			}

			// SnakeMovement

			for i := 0; i < counterTail; i++ {
				snakePosition[i] = snake[i].Position
			}

			if frameCounter%5 == 0 {
				for i := 0; i < counterTail; i++ {
					if i == 0 {
						snake[0].Position.X += snake[0].Speed.X
						snake[0].Position.Y += snake[0].Speed.Y
						allowMove = true
					} else {
						snake[i].Position = snakePosition[i-1]
					}
				}
			}

			// Wall behavior

			if (snake[0].Position.X > screenWidth-offset.X) || (snake[0].Position.Y > screenHeight-offset.Y) || (snake[0].Position.X < 0) || (snake[0].Position.Y < 0) {
				gameOver = true
			}

			// Collision with yourself

			for i := 0; i < counterTail; i++ {
				if (snake[0].Position.X == snake[i].Position.X) && (snake[0].Position.Y == snake[i].Position.X) {
					gameOver = true
				}
			}

			// Fruit Position calculation

			if !fruit.IsActive {
				fruit.IsActive = true
				fruit.Position = rl.Vector2{
					X: float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize) + offset.X/2,
					Y: float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize) + offset.Y/2,
				}

				// Ensure fruit does not overlap with the snake
				for i := 0; i < counterTail; i++ {
					for fruit.Position.X == snake[i].Position.X && fruit.Position.Y == snake[i].Position.Y {
						fruit.Position = rl.Vector2{
							X: float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize) + offset.X/2,
							Y: float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize) + offset.Y/2,
						}
						i = 0
					}
				}
			}

			// Collision detection between the snake's head and the fruit
			if (snake[0].Position.X < fruit.Position.X+fruit.Size.X && snake[0].Position.X+snake[0].Size.X > fruit.Position.X) &&
				(snake[0].Position.Y < fruit.Position.Y+fruit.Size.Y && snake[0].Position.Y+snake[0].Size.Y > fruit.Position.Y) {

				// Add a new segment to the snake's tail
				snake[counterTail].Position = snakePosition[counterTail-1]
				counterTail++
				score++
				fruit.IsActive = false
			}

			displayScore(score)
			frameCounter++
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			InitGame()
			gameOver = false
		}
	}
}

func DrawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	if !gameOver {
		// Drawgrid lines
		for i := 0; i < screenWidth/squareSize+1; i++ {
			rl.DrawLineV(
				rl.Vector2{X: float32(squareSize*i + int(offset.X)/2), Y: offset.Y / 2}, rl.Vector2{
					X: float32(squareSize*i + int(offset.X)/2),
					Y: screenHeight - offset.Y/2,
				}, rl.LightGray)
		}

		for i := 0; i < screenHeight/squareSize+1; i++ {
			rl.DrawLineV(rl.Vector2{X: offset.X / 2, Y: float32(squareSize*i + int(offset.Y)/2)}, rl.Vector2{X: screenWidth - offset.X/2, Y: float32(squareSize*i + int(offset.Y)/2)}, rl.LightGray)
		}

		// DrawSnake

		for i := 0; i < counterTail; i++ {
			rl.DrawRectangleV(snake[i].Position, snake[i].Size, snake[i].Color)

			// draw fruit to pick
			rl.DrawRectangleV(fruit.Position, fruit.Size, fruit.Color)

			if pause && !isStartMenu {
				rl.DrawText("Game Paused", screenWidth/2-rl.MeasureText("Game Paused", 40)/2, screenHeight/2-40, 40, rl.Gray)
			}

			if isStartMenu {
				startMenu()
			}
		}
	} else {
		rl.DrawText("Press [ENTER] to Play Again!", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [ENTER] to Play Again!", 20)/2, int32(rl.GetScreenHeight())/2-50, 20, rl.Gray)
		score = 0
	}

	rl.EndDrawing()
}

func UnloadGame() {
}

func UpdateDrawFrame() {
	UpdateGame()
	DrawGame()
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "SnakeGo")

	InitGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// rl.BeginDrawing()

		if isStartMenu {
			fmt.Println("We're on the start menu now")
		}
		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("Golang na scam: Shadow go to sleep!", 190, 200, 20, rl.Black)

		// rl.EndDrawing()

		UpdateDrawFrame()
	}
	UnloadGame()
	rl.CloseWindow()
}

func startMenu() {
	startText := "Start Menu"
	var fontSize int32 = 40
	rl.DrawText("Start Menu", screenWidth/2-rl.MeasureText(startText, fontSize)/2, screenHeight/2-fontSize, fontSize, rl.Gray)

	if rl.IsKeyPressed(rl.KeyEscape) {
		isStartMenu = false
	}
}
