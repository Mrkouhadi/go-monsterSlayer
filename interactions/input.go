package interactions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for { // or   for true {}   ... we keep asking for the player choice until we get what we want
		PlayerChoice, _ := getInputData()
		if PlayerChoice == "1" {
			return "ATTACK"
		} else if PlayerChoice == "2" {
			return "HEAL"
		} else if PlayerChoice == "3" && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Could not fetch the player's choice. Please try again !")
	}
}

func getInputData() (string, error) {
	fmt.Print("Your Choice: ") // text printed in the same line where the following will be printed
	inputData, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	inputData = strings.Replace(inputData, "\n", "", -1)

	return inputData, nil
}
