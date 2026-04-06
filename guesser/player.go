package guesser

import (
	"fmt"
)

type player struct {
	name string
	score int
	history []int
	difficulty int
}

func (pPlayer *player)ask_name () {
    fmt.Print("What is your name?\n")
	fmt.Scan(&pPlayer.name)
}
func Init_player(difficulty int) (new_player player) {
	new_player.ask_name()
	new_player.history = make([]int, 0)
	new_player.score = 0
	new_player.difficulty = difficulty
	return
}

func print_history(history []int) () {
	fmt.Printf("Already played:")
	for i := range history {
		fmt.Printf("%d ", history[i])
	}
	fmt.Print("\n")
}

func (pPlayer player)print_player () {
	fmt.Printf("%s, your current score is: %d.\n", pPlayer.name, pPlayer.score)
	print_history(pPlayer.history)
}