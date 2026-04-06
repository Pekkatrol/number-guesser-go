package guesser

import (
	"os"
	"fmt"
)

func Print_error(text string) {
    fmt.Fprint(os.Stderr, text)
    os.Exit(2)
}