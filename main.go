package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func main() {
	game := NewGame()
	game.RenderBoard()

	// Create a raw terminal that does not depend on enter for input
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	// Restore the old terminal when application is finished
	// Shell behaves strangely if not restored
	// If you use os.Exit() anywhere, defer is not called and
	// a manual cleanup is required
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	input := make([]byte, 3)

	for {
		n, err := os.Stdin.Read(input)

		if err != nil {
			break
		}
		// q or ctr+c cancel game
		if input[0] == 'q' || input[0] == 3 {
			break
		}

		if n == 3 && input[0] == 0x1b && input[1] == '[' {
			switch input[2] {
			case 'A':
				fmt.Println("up")

			case 'B':
				fmt.Println("down")

			case 'C':
				fmt.Println("right")

			case 'D':
				fmt.Println("left")
			}
		}
	}
}
