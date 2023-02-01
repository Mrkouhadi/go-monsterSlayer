package main

import (
	"github.com/mrkouhadi/go-monsterSlayer/interactions"
)

var CurrentRound = 0

func main() {
	StartGame()
	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = ExecuteRound()
	}
	Endgame()
}

func StartGame() {
	interactions.ShowGreeting()
}

func ExecuteRound() string { // will return "Player" || "Monster" || "" for winner variable
	CurrentRound++
	isSpecialRound := CurrentRound%3 == 0 // if CurrentRound%3 == 0 then isSpecialRound = true
	interactions.ShowAvailableActions(isSpecialRound)
	userChoice := interactions.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {

	} else if userChoice == "HEAL" {

	} else {

	}
	return ""
}

func Endgame() {

}
