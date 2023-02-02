package main

import (
	"github.com/mrkouhadi/go-monsterSlayer/actions"
	"github.com/mrkouhadi/go-monsterSlayer/interactions"
)

var CurrentRound = 0

func main() {
	StartGame()
	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = ExecuteRound()
	}
	Endgame(winner)
}

func StartGame() {
	interactions.ShowGreeting()
}

func ExecuteRound() string { // will return "Player" || "Monster" || "" for winner variable
	CurrentRound++
	isSpecialRound := CurrentRound%3 == 0 // if CurrentRound%3 == 0 then isSpecialRound = true
	interactions.ShowAvailableActions(isSpecialRound)
	userChoice := interactions.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}
	actions.AttackPlayer()
	playerHealth, monsterHealth = actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}

func Endgame(w string) {
	interactions.DeclareWinner(w)
}
