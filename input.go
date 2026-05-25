package main

import (
	"os"
)

func readInput(ch chan<- string) error {

	buf := make([]byte, 3)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		return err
	}

	// q or ctr+c cancel game
	if buf[0] == 'q' || buf[0] == 3 {
		ch <- "quit"
		return nil
	}

	var input string
	if n == 3 && buf[0] == 0x1b && buf[1] == '[' {
		switch buf[2] {
		case 'A':
			input = "up"

		case 'B':
			input = "down"
		case 'C':
			input = "right"
		case 'D':
			input = "left"

		default:
			input = ""
		}
	}
	ch <- input

	return nil
}
