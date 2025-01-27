package main

import (
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
	counterTail  int  = 0
	offset       rl.Vector2
)

type (
	SnakeSegment struct {
		position rl.Vector2
		size     rl.Vector2
		speed    rl.Vector2
		color    rl.Color
	}

	Food struct {
		position rl.Vector2
		size     rl.Vector2
		isActive bool
		color    rl.Color
	}
)

var (
	fruit         Food
	snake         []SnakeSegment = make([]SnakeSegment, snakeLength)
	snakePosition []rl.Vector2   = make([]rl.Vector2, snakeLength)
)

func InitGame() {
	frameCounter = 0
	gameOver = false
	pause = false

	counterTail = 1
	allowMove = false

	offset.X = screenWidth % squareSize
	offset.Y = screenHeight % squareSize

	for i := 0; i < snakeLength; i++ {
		snake[i].position = rl.Vector2{X: offset.X / 2, Y: offset.Y / 2}
		snake[i].size = rl.Vector2{X: squareSize, Y: squareSize}
		snake[i].speed = rl.Vector2{X: squareSize, Y: 0}

		if i == 0 {
			snake[i].color = rl.DarkBlue
		} else {
			snake[i].color = rl.Blue
		}
	}

	for i := 0; i < snakeLength; i++ {
		snakePosition[i] = rl.Vector2{X: 0.0, Y: 0.0}
	}

	fruit.size = rl.Vector2{X: squareSize, Y: squareSize}
	fruit.color = rl.SkyBlue
	fruit.isActive = false
}

func UpdateGame() {
	if !gameOver {
		if rl.IsKeyPressed('P') {
			pause = !pause
		}

		if !pause {
			// Player controls
			if rl.IsKeyPressed(rl.KeyRight) && snake[0].speed.X == 0 && allowMove {
				snake[0].speed = rl.Vector2{X: squareSize, Y: 0}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyLeft) && snake[0].speed.X == 0 && allowMove {
				snake[0].speed = rl.Vector2{X: -squareSize, Y: 0}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyUp) && snake[0].speed.Y == 0 && allowMove {
				snake[0].speed = rl.Vector2{X: 0, Y: -squareSize}
				allowMove = false
			}

			if rl.IsKeyPressed(rl.KeyDown) && snake[0].speed.Y == 0 && allowMove {
				snake[0].speed = rl.Vector2{X: 0, Y: squareSize}
				allowMove = false
			}

			// SnakeMovement

			for i := 0; i < counterTail; i++ {
				snakePosition[i] = snake[i].position
			}

			if frameCounter%5 == 0 {
				for i := 0; i < counterTail; i++ {
					if i == 0 {
						snake[0].position.X += snake[0].speed.X
						snake[0].position.Y += snake[0].speed.Y
						allowMove = true
					} else {
						snake[i].position = snakePosition[i-1]
					}
				}
			}

			// Wall behavior

			if (snake[0].position.X > screenWidth-offset.X) || (snake[0].position.Y > screenHeight-offset.Y) || (snake[0].position.X < 0) || (snake[0].position.Y < 0) {
				gameOver = true
			}

			// Collision with yourself

			for i := 0; i < counterTail; i++ {
				if (snake[0].position.X == snake[i].position.X) && (snake[0].position.Y == snake[i].position.X) {
					gameOver = true
				}
			}

			// Fruit Position calculation

			if !fruit.isActive {
				fruit.isActive = true
				fruit.position = rl.Vector2{
					X: float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize) + offset.X/2,
					Y: float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize) + offset.Y/2,
				}
			
				// Ensure fruit does not overlap with the snake
				for i := 0; i < counterTail; i++ {
					for fruit.position.X == snake[i].position.X && fruit.position.Y == snake[i].position.Y {
						fruit.position = rl.Vector2{
							X: float32(rl.GetRandomValue(0, (screenWidth/squareSize)-1)*squareSize) + offset.X/2,
							Y: float32(rl.GetRandomValue(0, (screenHeight/squareSize)-1)*squareSize) + offset.Y/2,
						}
						i = 0
					}
				}
			}
			
			// Collision detection between the snake's head and the fruit
			if (snake[0].position.X < fruit.position.X+fruit.size.X && snake[0].position.X+snake[0].size.X > fruit.position.X) &&
				(snake[0].position.Y < fruit.position.Y+fruit.size.Y && snake[0].position.Y+snake[0].size.Y > fruit.position.Y) {
			
				// Add a new segment to the snake's tail
				snake[counterTail].position = snakePosition[counterTail-1]
				counterTail++
				fruit.isActive = false
			}
			

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
		//Drawgrid lines
		for i := 0; i < screenWidth/squareSize+1; i++ {
			rl.DrawLineV(
				rl.Vector2{X: float32(squareSize*i + int(offset.X)/2), Y: offset.Y / 2}, rl.Vector2{
					X: float32(squareSize*i + int(offset.X)/2),
					Y: screenHeight - offset.Y/2,
				}, rl.LightGray)
		}

		for i := 0; i<screenHeight/squareSize + 1; i++{
			rl.DrawLineV(rl.Vector2{X: offset.X/2, Y: float32(squareSize*i + int(offset.Y)/2)}, rl.Vector2{X: screenWidth - offset.X/2, Y: float32(squareSize*i + int(offset.Y)/2)}, rl.LightGray);
		}


		//DrawSnake

		for i := 0; i < counterTail; i++{
			rl.DrawRectangleV(snake[i].position, snake[i].size, snake[i].color)

			// draw fruit to pick
			rl.DrawRectangleV(fruit.position, fruit.size, fruit.color)

			if pause{
				rl.DrawText("Game Paused", screenWidth/2 - rl.MeasureText("Game Paused", 40)/2, screenHeight/2-40, 40, rl.Gray)
			}
		}
	}else{
		rl.DrawText("Press [ENTER] to Play Again!", int32(rl.GetScreenWidth())/2 - rl.MeasureText("Press [ENTER] to Play Again!", 20)/2, int32(rl.GetScreenHeight())/2 - 50, 20, rl.Gray)
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

		// rl.ClearBackground(rl.RayWhite)
		// rl.DrawText("Golang na scam: Shadow go to sleep!", 190, 200, 20, rl.Black)

		// rl.EndDrawing()

		UpdateDrawFrame()
	}
	UnloadGame()
	rl.CloseWindow()
}
