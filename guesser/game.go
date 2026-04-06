package guesser

import (
	"fmt"
	"math/rand"
)

const (
	EASY = iota + 1
	MEDIUM
	HARD
	IMPOSSIBLE
)

func ask_difficulty() (difficulty int) {
	fmt.Printf("Choose a difficulty:\n")
	fmt.Printf("1. Easy [0 - 20]\n")
	fmt.Printf("2. Medium [0 - 200]\n")
	fmt.Printf("3. Hard [0 - 2000]\n")
	fmt.Printf("4. Impossible [0 - 20000]\n")
	fmt.Scan(&difficulty)

	if (difficulty > 4 || difficulty < 1) {
		Print_error("Wrong choice\n")
	}
	return
}

func generate_number(difficulty int) (number int) {
	switch (difficulty) {
		case EASY:
			number = rand.Intn(20)
		case MEDIUM:
			number = rand.Intn(200)
		case HARD:
			number = rand.Intn(2000)
		case IMPOSSIBLE:
			number = rand.Intn(20000)
	}
	return
}

func print_hint(current int, number int) {
	if (current < number) {
		fmt.Printf("Try upper\n")
	} else if (current > number) {
		fmt.Printf("Try lower\n")
	}
}

func Start_game() () {
	difficulty := ask_difficulty()
	pPlayer := Init_player(difficulty)

	for {
		current := 0
		number := generate_number(pPlayer.difficulty)
		score := 10
		for {
			pPlayer.print_player()
			fmt.Printf("Choose a number: ")
			fmt.Scan(&current)
			pPlayer.history = append(pPlayer.history, current)
			if (number == current) {
				if (score > 0) {
				    pPlayer.score += score * pPlayer.difficulty
				}
				break
			}
			score--
			print_hint(current, number)
		}
		fmt.Printf("Your score is: %d\n", pPlayer.score)
		pPlayer.history = make([]int, 0)
	}
}