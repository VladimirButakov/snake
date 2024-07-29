package entities

import "math/rand"

type Food struct {
	Position                            Vector2Int
	screenWidth, screenHeight, gridSize int
}

func NewFood(screenWidth, screenHeight, gridSize int) *Food {
	return &Food{
		Position:     Vector2Int{X: rand.Intn(screenWidth / gridSize), Y: rand.Intn(screenHeight / gridSize)},
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		gridSize:     gridSize,
	}
}

func (f *Food) Respawn() {
	f.Position = Vector2Int{X: rand.Intn(f.screenWidth / f.gridSize), Y: rand.Intn(f.screenHeight / f.gridSize)}
}
