package frameworks

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/VladimirButakov/snake/internal/usecases"
)

type RaylibAdapter struct {
	screenWidth, screenHeight, gridSize int
}

func NewRaylibAdapter(screenWidth, screenHeight, gridSize int) *RaylibAdapter {
	return &RaylibAdapter{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		gridSize:     gridSize,
	}
}

func (r *RaylibAdapter) InitWindow() {
	rl.InitWindow(int32(r.screenWidth), int32(r.screenHeight), "Snake Game")
}

func (r *RaylibAdapter) CloseWindow() {
	rl.CloseWindow()
}

func (r *RaylibAdapter) SetTargetFPS(fps int) {
	rl.SetTargetFPS(int32(fps))
}

func (r *RaylibAdapter) WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

func (r *RaylibAdapter) GetTime() float32 {
	return float32(rl.GetTime())
}

func (r *RaylibAdapter) IsKeyPressed(key int) bool {
	return rl.IsKeyPressed(int32(key))
}

func (r *RaylibAdapter) BeginDrawing() {
	rl.BeginDrawing()
}

func (r *RaylibAdapter) ClearBackground() {
	rl.ClearBackground(rl.RayWhite)
}

func (r *RaylibAdapter) DrawRectangle(x, y, width, height int, color usecases.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), rl.NewColor(color.R, color.G, color.B, color.A))
}

func (r *RaylibAdapter) DrawText(text string, x, y, fontSize int, color usecases.Color) {
	rl.DrawText(text, int32(x), int32(y), int32(fontSize), rl.NewColor(color.R, color.G, color.B, color.A))
}

func (r *RaylibAdapter) EndDrawing() {
	rl.EndDrawing()
}

func (r *RaylibAdapter) GetKeyRight() int {
	return rl.KeyRight
}

func (r *RaylibAdapter) GetKeyLeft() int {
	return rl.KeyLeft
}

func (r *RaylibAdapter) GetKeyUp() int {
	return rl.KeyUp
}

func (r *RaylibAdapter) GetKeyDown() int {
	return rl.KeyDown
}

func (r *RaylibAdapter) GetKeyRestart() int {
	return rl.KeySpace
}
