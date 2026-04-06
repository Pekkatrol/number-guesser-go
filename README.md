# Number Guesser (Go)

A simple command-line number guessing game written in Go.

The game asks for a difficulty level and your player name, then generates a random number to guess.  
After each guess, you get a hint (`Try upper` / `Try lower`) until you find the correct number.

---

## Features

- Interactive CLI game
- 4 difficulty levels:
  - **Easy**: number between `0` and `19`
  - **Medium**: number between `0` and `199`
  - **Hard**: number between `0` and `1999`
  - **Impossible**: number between `0` and `19999`
- Score system:
  - Starts at `10` for each round
  - Decreases after wrong guesses
  - Round points are multiplied by difficulty
- Guess history displayed during the round

---

## Requirements

- Go installed (tested with Go 1.22+)

If Go is not installed on Ubuntu/Debian:

```bash
sudo apt update
sudo apt install golang-go
```

---

## How to run

From the project root:

```bash
go run .
```

> The program expects **no arguments**.

---

## How it works

1. Choose a difficulty (`1` to `4`)
2. Enter your name
3. Guess the generated number
4. Read hints after each wrong attempt
5. Score is updated when you win the round
6. A new round starts automatically

---

## Example

```text
Choose a difficulty:
1. Easy [0 - 20]
2. Medium [0 - 200]
3. Hard [0 - 2000]
4. Impossible [0 - 20000]
1
What is your name?
Alex
Alex, your current score is: 0.
Already played:
Choose a number: 10
Try lower
Alex, your current score is: 0.
Already played:10
Choose a number: 7
Your score is: 9
```

---

## Error handling

The program exits with an error if:

- arguments are provided
- difficulty choice is invalid

---

## Notes

- Random number generation uses Go's `math/rand`.
- The game currently runs in an infinite loop (stop with `Ctrl + C`).

Enjoy the game!
