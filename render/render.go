package render

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SnakeLength = 256
	SquareSize  = 31

	ScreenWidth  = 800
	ScreenHeight = 500
)

var (
	IsGameOver     bool = false
	FrameCounter int  = 0
	IsGamePaused       bool = false
	AllowMove    bool = false
	IsStartMenu  bool = false
	Score        int  = 0
	CounterTail  int  = 0
	Offset       rl.Vector2
)


var (
	fruit        Food
	snake         []SnakeSegment = make([]SnakeSegment, SnakeLength)
	snakePosition []rl.Vector2   = make([]rl.Vector2, SnakeLength)
)

func InitGame() {
	FrameCounter = 0
	IsGameOver = false
	IsGamePaused = false
	IsStartMenu = false

	CounterTail = 1
	AllowMove = false

	Offset.X = ScreenWidth % SquareSize
	Offset.Y = ScreenHeight % SquareSize

	for i := 0; i < SnakeLength; i++ {
		snake[i].Position = rl.Vector2{X: Offset.X / 2, Y: Offset.Y / 2}
		snake[i].Size = rl.Vector2{X: SquareSize, Y: SquareSize}
		snake[i].Speed = rl.Vector2{X: SquareSize, Y: 0}

		if i == 0 {
			snake[i].Color = rl.DarkBlue
		} else {
			snake[i].Color = rl.Blue
		}
	}

	for i := 0; i < SnakeLength; i++ {
		snakePosition[i] = rl.Vector2{X: 0.0, Y: 0.0}
	}

	fruit.Size = rl.Vector2{X: SquareSize, Y: SquareSize}
	fruit.Color = rl.SkyBlue
	fruit.IsActive = false
	// displayScore(score)
}

func UpdateGame() {
	if !IsGameOver {
		if rl.IsKeyPressed('P') {
			IsGamePaused = !IsGamePaused
		}

		if rl.IsKeyPressed(rl.KeyM) {
			IsStartMenu = !IsStartMenu
		}

		if !IsGamePaused && !IsStartMenu {
			// Player controls
			if rl.IsKeyPressed(rl.KeyRight) && snake[0].Speed.X == 0 && AllowMove {
				snake[0].Speed = rl.Vector2{X: SquareSize, Y: 0}
				AllowMove = false
			}

			if rl.IsKeyPressed(rl.KeyLeft) && snake[0].Speed.X == 0 && AllowMove {
				snake[0].Speed = rl.Vector2{X: -SquareSize, Y: 0}
				AllowMove = false
			}

			if rl.IsKeyPressed(rl.KeyUp) && snake[0].Speed.Y == 0 && AllowMove {
				snake[0].Speed = rl.Vector2{X: 0, Y: -SquareSize}
				AllowMove = false
			}

			if rl.IsKeyPressed(rl.KeyDown) && snake[0].Speed.Y == 0 && AllowMove {
				snake[0].Speed = rl.Vector2{X: 0, Y: SquareSize}
				AllowMove = false
			}

			// SnakeMovement

			for i := 0; i < CounterTail; i++ {
				snakePosition[i] = snake[i].Position
			}

			if FrameCounter%5 == 0 {
				for i := 0; i < CounterTail; i++ {
					if i == 0 {
						snake[0].Position.X += snake[0].Speed.X
						snake[0].Position.Y += snake[0].Speed.Y
						AllowMove = true
					} else {
						snake[i].Position = snakePosition[i-1]
					}
				}
			}

			// Wall behavior

			if (snake[0].Position.X > ScreenWidth-Offset.X) || (snake[0].Position.Y > ScreenHeight-Offset.Y) || (snake[0].Position.X < 0) || (snake[0].Position.Y < 0) {
				IsGameOver = true
			}

			// Collision with yourself

			for i := 0; i < CounterTail; i++ {
				if (snake[0].Position.X == snake[i].Position.X) && (snake[0].Position.Y == snake[i].Position.X) {
					IsGameOver = true
				}
			}

			// Fruit Position calculation

			if !fruit.IsActive {
				fruit.IsActive = true
				fruit.Position = rl.Vector2{
					X: float32(rl.GetRandomValue(0, (ScreenWidth/SquareSize)-1)*SquareSize) + Offset.X/2,
					Y: float32(rl.GetRandomValue(0, (ScreenHeight/SquareSize)-1)*SquareSize) + Offset.Y/2,
				}

				// Ensure fruit does not overlap with the snake
				for i := 0; i < CounterTail; i++ {
					for fruit.Position.X == snake[i].Position.X && fruit.Position.Y == snake[i].Position.Y {
						fruit.Position = rl.Vector2{
							X: float32(rl.GetRandomValue(0, (ScreenWidth/SquareSize)-1)*SquareSize) + Offset.X/2,
							Y: float32(rl.GetRandomValue(0, (ScreenHeight/SquareSize)-1)*SquareSize) + Offset.Y/2,
						}
						i = 0
					}
				}
			}

			// Collision detection between the snake's head and the fruit
			if (snake[0].Position.X < fruit.Position.X+fruit.Size.X && snake[0].Position.X+snake[0].Size.X > fruit.Position.X) &&
				(snake[0].Position.Y < fruit.Position.Y+fruit.Size.Y && snake[0].Position.Y+snake[0].Size.Y > fruit.Position.Y) {

				// Add a new segment to the snake's tail
				snake[CounterTail].Position = snakePosition[CounterTail-1]
				CounterTail++
				Score++
				fruit.IsActive = false
			}

			
			FrameCounter++
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			InitGame()
			IsGameOver = false
		}
	}
}

func DrawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	var gameObjects []Drawable

	if !IsGameOver && !IsStartMenu{

		// Drawgrid lines
		for i := 0; i < ScreenWidth/SquareSize+1; i++ {
			rl.DrawLineV(
				rl.Vector2{X: float32(SquareSize*i + int(Offset.X)/2), Y: Offset.Y / 2}, rl.Vector2{
					X: float32(SquareSize*i + int(Offset.X)/2),
					Y: ScreenHeight - Offset.Y/2,
				}, rl.LightGray)
		}

		for i := 0; i < ScreenHeight/SquareSize+1; i++ {
			rl.DrawLineV(rl.Vector2{X: Offset.X / 2, Y: float32(SquareSize*i + int(Offset.Y)/2)}, rl.Vector2{X: ScreenWidth - Offset.X/2, Y: float32(SquareSize*i + int(Offset.Y)/2)}, rl.LightGray)
		}

		// DrawSnake

		for i := 0; i < CounterTail; i++ {

			newSnake := &SnakeSegment{Position: snake[i].Position, Size: snake[i].Size, Color: snake[i].Color, Zindex: 1}
			newFruit := &Food{Position: fruit.Position, Size: fruit.Size, Color: fruit.Color, Zindex: 1}

			newGameScore := GameScoreUI(Score)
		

			gameObjects = append(gameObjects, newSnake)
			gameObjects = append(gameObjects, newFruit)
			gameObjects = append(gameObjects, newGameScore)

			if IsGamePaused && !IsStartMenu {
				newGameIsGamePausedUI := GamePausedUI(ScreenWidth, ScreenHeight)
				gameObjects = append(gameObjects, newGameIsGamePausedUI)
			}

			if IsStartMenu {
				startMenu()
			}
		}
	} else {
		rl.DrawText("Press [ENTER] to Play Again!", int32(rl.GetScreenWidth())/2-rl.MeasureText("Press [ENTER] to Play Again!", 20)/2, int32(rl.GetScreenHeight())/2-50, 20, rl.Gray)
		Score = 0
	}
	

	// Sort objects according to their ZIndex
	sort.Slice(gameObjects, func(i, j int)bool{
		return gameObjects[i].ZIndex() < gameObjects[j].ZIndex()
	})

	//Render objects inorder of zindex
	for _, obj := range gameObjects{
		obj.Draw()
	}

	rl.EndDrawing()
}

func UnloadGame() {
}

func UpdateDrawFrame() {
	UpdateGame()
	DrawGame()
}

func startMenu() {
	startText := "Start Menu"
	var fontSize int32 = 40
	rl.DrawText("Start Menu", ScreenWidth/2-rl.MeasureText(startText, fontSize)/2, ScreenHeight/2-fontSize, fontSize, rl.Gray)
}
