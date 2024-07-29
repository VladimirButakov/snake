package entities

type Snake struct {
	Body      []Vector2Int
	Direction Vector2Int
	Grow      bool
	Score     int
}

func NewSnake(x, y int) *Snake {
	return &Snake{
		Body:      []Vector2Int{{X: x, Y: y}},
		Direction: Vector2Int{X: 1, Y: 0},
		Grow:      false,
		Score:     0,
	}
}

func (s *Snake) SetDirection(x, y int) {
	if s.Direction.X != -x && s.Direction.Y != -y {
		s.Direction = Vector2Int{X: x, Y: y}
	}
}

func (s *Snake) Move(screenWidth, screenHeight, gridSize int) bool {
	head := s.Body[0]
	newHead := Vector2Int{X: head.X + s.Direction.X, Y: head.Y + s.Direction.Y}

	if newHead.X < 0 {
		newHead.X = screenWidth/gridSize - 1
	} else if newHead.X >= screenWidth/gridSize {
		newHead.X = 0
	}
	if newHead.Y < 0 {
		newHead.Y = screenHeight/gridSize - 1
	} else if newHead.Y >= screenHeight/gridSize {
		newHead.Y = 0
	}

	for _, segment := range s.Body[1:] {
		if segment == newHead {
			return true
		}
	}

	s.Body = append([]Vector2Int{newHead}, s.Body...)

	if !s.Grow {
		s.Body = s.Body[:len(s.Body)-1]
	} else {
		s.Grow = false
	}

	return false
}

func (s *Snake) CheckFood(food *Food) bool {
	if s.Body[0] == food.Position {
		s.Grow = true
		s.Score++ // Увеличиваем счет
		return true
	}
	return false
}
