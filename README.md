# Snake Game in Go

This is a simple implementation of the classic Snake game written in Go using the Raylib library. The game features basic mechanics including snake movement, food consumption, collision detection, and scoring.

## Features

- Classic Snake game mechanics
- Snake grows longer when eating food
- Game ends when the snake collides with itself
- Score display
- Game restart functionality with spacebar

## Prerequisites

To run this game, you'll need:

- [Go](https://golang.org/doc/install) installed on your machine
- [Raylib-Go](https://github.com/gen2brain/raylib-go) library

## Installation

1. **Clone the repository:**

   ```
   bash
   git clone https://github.com/VladimirButakov/snake.git
   cd snake
   ```
   
2. **Build and run the game:**

   ``` go run cmd/snake/main.go ```

## Game Controls

**Arrow Keys: Control the direction of the snake.**

**Spacebar: Restart the game after a game over.**