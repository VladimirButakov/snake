package adapters

import "github.com/VladimirButakov/snake/internal/usecases"

func RunGame(game *usecases.Game) {
	game.Run()
}
