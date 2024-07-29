package main

import (
	"github.com/VladimirButakov/snake/internal/adapters"
	"github.com/VladimirButakov/snake/internal/frameworks"
	"github.com/VladimirButakov/snake/internal/usecases"
)

const (
	screenWidth  = 800
	screenHeight = 450
	gridSize     = 20
	snakeSpeed   = 0.1
)

func main() {
	raylibAdapter := frameworks.NewRaylibAdapter(screenWidth, screenHeight, gridSize)
	game := usecases.NewGame(screenWidth, screenHeight, gridSize, snakeSpeed, raylibAdapter)
	adapters.RunGame(game)
}
