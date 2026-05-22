package main

import (
	"fmt"
	"strings"
)

type game struct {
	board [][]string
}

type point struct {
	x int
	y int
}

func NewGame() game {
	// Methode 1: make + Loop (empfohlen, zusammenhängender Speicher)
	matrix := createMatrix()

	return game{matrix}
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

func (g *game) RenderBoard() {
	for _, row := range g.board {
		fmt.Println(strings.Join(row, ""))
	}
}
