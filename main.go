package main

import (
	"github.com/mrkouhadi/go-monsterSlayer/actions"
	"github.com/mrkouhadi/go-monsterSlayer/interactions"
)

var CurrentRound = 0
var Rounds = []interactions.RoundData{}

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

	var playerAttackdmg int
	var playerhealVal int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackdmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerhealVal = actions.HealPlayer()
	} else {
		playerAttackdmg = actions.AttackMonster(true)
	}
	monsterAttackDmg = actions.AttackPlayer()
	playerHealth, monsterHealth := actions.GetHealthAmounts()

	rounData := interactions.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackdmg,
		PlayerHealValue:  playerhealVal,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interactions.ShowRoundStatistics(&rounData)

	Rounds = append(Rounds, rounData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func Endgame(w string) {
	interactions.DeclareWinner(w)
	interactions.WriteDataIntoFile(&Rounds)
}
