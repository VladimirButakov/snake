package usecases

import (
	"strconv"

	"github.com/VladimirButakov/snake/internal/entities"
)

type Game struct {
	screenWidth    int
	screenHeight   int
	gridSize       int
	snakeSpeed     float32
	snake          *entities.Snake
	food           *entities.Food
	lastUpdateTime float32
	adapter        GameAdapter
	isGameOver     bool
}

type GameAdapter interface {
	InitWindow()
	CloseWindow()
	SetTargetFPS(int)
	WindowShouldClose() bool
	GetTime() float32
	IsKeyPressed(int) bool
	BeginDrawing()
	ClearBackground()
	DrawRectangle(int, int, int, int, Color)
	DrawText(string, int, int, int, Color)
	EndDrawing()
	GetKeyRight() int
	GetKeyLeft() int
	GetKeyUp() int
	GetKeyDown() int
	GetKeyRestart() int
}

type Color struct {
	R, G, B, A uint8
}

func NewGame(screenWidth, screenHeight, gridSize int, snakeSpeed float32, adapter GameAdapter) *Game {
	food := entities.NewFood(screenWidth, screenHeight, gridSize)
	snake := entities.NewSnake(10, 10)

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		gridSize:     gridSize,
		snakeSpeed:   snakeSpeed,
		snake:        snake,
		food:         food,
		adapter:      adapter,
		isGameOver:   false,
	}
}

func (g *Game) Run() {
	g.adapter.InitWindow()
	defer g.adapter.CloseWindow()
	g.adapter.SetTargetFPS(60)

	for !g.adapter.WindowShouldClose() {
		currentTime := g.adapter.GetTime()

		if g.isGameOver {
			if g.adapter.IsKeyPressed(g.adapter.GetKeyRestart()) {
				g.Reset()
			}
			g.adapter.BeginDrawing()
			g.adapter.ClearBackground()
			g.adapter.DrawText("Game Over! Press SPACE to Restart", g.screenWidth/2-150, g.screenHeight/2, 20, ColorRed())
			g.adapter.EndDrawing()
			continue
		}

		if g.adapter.IsKeyPressed(g.adapter.GetKeyRight()) {
			g.snake.SetDirection(1, 0)
		} else if g.adapter.IsKeyPressed(g.adapter.GetKeyLeft()) {
			g.snake.SetDirection(-1, 0)
		} else if g.adapter.IsKeyPressed(g.adapter.GetKeyUp()) {
			g.snake.SetDirection(0, -1)
		} else if g.adapter.IsKeyPressed(g.adapter.GetKeyDown()) {
			g.snake.SetDirection(0, 1)
		}

		if currentTime-g.lastUpdateTime >= g.snakeSpeed {
			if g.snake.Move(g.screenWidth, g.screenHeight, g.gridSize) {
				g.isGameOver = true
			} else if g.snake.CheckFood(g.food) {
				g.food.Respawn()
			}
			g.lastUpdateTime = currentTime
		}

		g.adapter.BeginDrawing()
		g.adapter.ClearBackground()

		for _, segment := range g.snake.Body {
			g.adapter.DrawRectangle(segment.X*g.gridSize, segment.Y*g.gridSize, g.gridSize, g.gridSize, ColorGreen())
		}

		g.adapter.DrawRectangle(g.food.Position.X*g.gridSize, g.food.Position.Y*g.gridSize, g.gridSize, g.gridSize, ColorRed())

		g.adapter.DrawText("Score: "+strconv.Itoa(g.snake.Score), 10, 10, 20, ColorBlack())

		g.adapter.EndDrawing()
	}
}

func (g *Game) Reset() {
	g.snake = entities.NewSnake(10, 10)
	g.food.Respawn()
	g.lastUpdateTime = 0
	g.isGameOver = false
}

func ColorGreen() Color {
	return Color{R: 0, G: 255, B: 0, A: 255}
}

func ColorRed() Color {
	return Color{R: 255, G: 0, B: 0, A: 255}
}

func ColorBlack() Color {
	return Color{R: 0, G: 0, B: 0, A: 255}
}
