package interactions

import (
	"fmt"
	"os"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func ShowGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good Luck !")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action: ")
	fmt.Println("---------------------------")
	fmt.Println("Action (1) Attack Monster ")
	fmt.Println("Action (2) Heal Yourself ")

	if specialAttackIsAvailable {
		fmt.Println("Action (3) Special Attack ")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("---------------------------")
	fmt.Println("GAME OVER")
	fmt.Println("---------------------------")
	fmt.Printf("%v has won :D \n", winner)

}

func ShowRoundStatistics(Rdata *RoundData) {
	if Rdata.Action == "ATTACK" {
		fmt.Printf("Player Attacked Monster for %v damage. \n", Rdata.PlayerAttackDmg)
	} else if Rdata.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player hit with a strong Attack against Monster for %v damage. \n", Rdata.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", Rdata.PlayerHealValue)
	}

	fmt.Printf("Monster Attacked Player for %v damage. \n", Rdata.MonsterAttackDmg)
	fmt.Printf("Player's health: %v \n", Rdata.PlayerHealth)
	fmt.Printf("Monster's health: %v \n", Rdata.MonsterHealth)
}

func WriteDataIntoFile(rounds *[]RoundData) {
	file, err := os.Create("gameLog.txt")
	if err != nil {
		fmt.Println("Failed to Store Data in the log File.")
		return
	}
	for i, val := range *rounds {

		data := map[string]string{
			"Round":                 fmt.Sprint(i + 1),
			"Action":                val.Action,
			"Player Attack Damage":  fmt.Sprint(val.PlayerAttackDmg), // convert int to string
			"Player Heal Value":     fmt.Sprint(val.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(val.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(val.PlayerHealth),
			"Monster Health":        fmt.Sprint(val.MonsterHealth),
		}
		logEntry := fmt.Sprintln(data)
		_, err := file.WriteString(logEntry)
		if err != nil {
			fmt.Println("Failed to Store This round data to the LOG File.")
			continue
		}
	}
	file.Close()
	fmt.Println("Data of all round have been store to the LOG File.")
}
