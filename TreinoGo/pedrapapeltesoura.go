package main

import "fmt"

func jogar(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}

	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}

	return "Player 2 won!"
}

func main() {
	fmt.Println(jogar("rock", "scissors"))
}
