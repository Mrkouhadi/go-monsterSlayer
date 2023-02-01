package main

import "github.com/mrkouhadi/go-monsterSlayer/interactions"

func main() {
	StartGame()
	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = ExecuteRound()
	}
	Endgame()
}

func StartGame() {

}
func ExecuteRound() string { // will return "Player" || "Monster" || "" for winner variable
	interactions.PrintGreeting()
	return ""
}
func Endgame() {

}
