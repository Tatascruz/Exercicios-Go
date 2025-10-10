package main

import (
	"fmt"
	"strings"
)

func Points(games []string) int {
	totalPoints := 0
	for _, game := range games {
		score := strings.Split(game, ":")

		if score[0] == score[1] {
			totalPoints += 1
		}

		if score[0] > score[1] {
			totalPoints += 3
		}
	}
	return totalPoints

}

func main() {
	games := []string{"5:1", "4:4", "3:1", "0:1"}
	result := Points(games)
	fmt.Println("Total de pontos:", result)

}
