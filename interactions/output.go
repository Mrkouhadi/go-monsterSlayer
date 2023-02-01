package interactions

import "fmt"

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
