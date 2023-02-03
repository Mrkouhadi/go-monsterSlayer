package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currMonsterHealth = MONSTERER_HEALTH
var currPlayerHealth = PLAYER_HEALTH

func AttackMonster(IsSpecialAttack bool) int {
	minAttackVal := PLAYER_ATTACK_MIN_DMG
	maxAttackVal := PLAYER_ATTACK_MAX_DMG

	if IsSpecialAttack {
		minAttackVal = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackVal = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
	return dmgVal
}

func HealPlayer() int {

	healVal := generateRandBetween(PLAYER_HEAL_MIN_VAL, PLAYER_HEAL_MAX_VAL)
	healthDiff := PLAYER_HEALTH - currPlayerHealth

	if healthDiff >= healVal {
		currPlayerHealth += healVal
		return healVal
	} else {
		currPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}
}

func AttackPlayer() int {
	minAttackVal := MONSTER_ATTACK_MIN_DMG
	maxAttackVal := MONSTER_ATTACK_MAX_DMG

	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
	return dmgVal
}

func GetHealthAmounts() (int, int) {
	return currPlayerHealth, currMonsterHealth
}
func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
