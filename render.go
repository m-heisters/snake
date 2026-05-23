package main

import (
	"fmt"
	"strings"
)

func (g *game) RenderBoard() {
	for _, row := range g.board {
		fmt.Println(strings.Join(row, ""))
	}
}
