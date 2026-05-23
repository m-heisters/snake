package main

import (
	"math/rand"
)

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
	initialSnake := initSnakeArray()
	game := game{matrix, initialSnake}
	game.addSnakeToBoard(initialSnake)

	return game

}

func initSnakeArray() []point {
	x := rand.Intn(20) - 1
	y := rand.Intn(20) - 1
	snakeHead := point{
		x,
		y,
	}
	return []point{snakeHead}

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

func (g *game) addSnakeToBoard(points []point) {
	for _, point := range points {
		g.board[point.x][point.y] = "o"
	}

}

func (g *game) addPoint(p point) {

}
