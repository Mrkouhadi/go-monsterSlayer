package interactions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
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

	// disply " MONSTER SLAYER " using ASCII ART
	myfigure := figure.NewColorFigure("MONSTER SLAYER", "", "blue", true)
	myfigure.Print()

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

	// disply "Game over" using ASCII ART
	myfigure := figure.NewColorFigure("GAME OVER", "", "green", true)
	myfigure.Print()

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
	// when making an excutable program, we need to get the path where is the program is located
	expPath, err := os.Executable()
	if err != nil {
		fmt.Println("Could not get the path ofthe executable program ! ")
		return
	}
	expPath = filepath.Dir(expPath)
	file, err := os.Create(expPath + "/gameLog.txt")
	// to test this code (development mode) we need to comment out the code above and
	// uncomment the line of code below and then "go run main.go"
	// file, err := os.Create("gameLog.txt")

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
