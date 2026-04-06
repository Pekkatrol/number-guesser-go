package main

import (
	"os"
	"guesser/guesser"
)

func main () {
	args := os.Args
	
	if (len(args) != 1) {
		guesser.Print_error("No arguments needed\n")
	}
	guesser.Start_game()
}