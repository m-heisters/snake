package main

import (
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

}
