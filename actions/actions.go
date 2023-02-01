package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currMonsterHealth = 100
var currPlayerHealth = 100

func AttackMonster(IsSpecialAttack bool) {
	minAttackVal := 5
	maxAttackVal := 10

	if IsSpecialAttack {
		minAttackVal = 10
		maxAttackVal = 20
	}
	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
}

func HealPlayer() {
	minhealkVal := 10
	maxhealkVal := 20

	healVal := generateRandBetween(minhealkVal, maxhealkVal)
	healthDiff := 100 - currPlayerHealth

	if healthDiff >= healVal {
		currPlayerHealth += healVal
	} else {
		currPlayerHealth = 100
	}
}

func AttackPlayer() {
	minAttackVal := 8
	maxAttackVal := 15

	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
}

func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
