package main

type game struct {
	board [][]string
	snake []point
}

type point struct {
	x int
	y int
}

func NewGame() game {
	matrix := createMatrix()
	snakeHead := initSnakeArray()
	game := game{matrix, snakeHead}
	game.addSnakeToMatrix(snakeHead)

	return game

}

func initSnakeArray() []point {
	snake := []point{
		point{9, 9},
		point{9, 8},
		point{9, 7},
	}
	return snake

}

func createMatrix() [][]string {
	matrix := make([][]string, 20)
	for i := range matrix {
		matrix[i] = make([]string, 20)
		for j := range matrix[i] {
			if i == 0 || i == 19 || j == 0 || j == 19 {
				matrix[i][j] = "#"
			} else {
				matrix[i][j] = " "
			}
		}
	}
	return matrix
}

func (g *game) addSnakeToMatrix(snake []point) {
	for i, point := range snake {
		if i == 0 {
			g.board[point.x][point.y] = "@"

		} else {
			g.board[point.x][point.y] = "o"
		}
	}

}

func (g *game) addPoint(p point) {
	g.snake = append(g.snake, p)

}
