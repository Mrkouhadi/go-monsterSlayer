package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currMonsterHealth = MONSTERER_HEALTH
var currPlayerHealth = PLAYER_HEALTH

func AttackMonster(IsSpecialAttack bool) {
	minAttackVal := PLAYER_ATTACK_MIN_DMG
	maxAttackVal := PLAYER_ATTACK_MAX_DMG

	if IsSpecialAttack {
		minAttackVal = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackVal = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
}

func HealPlayer() {

	healVal := generateRandBetween(PLAYER_HEAL_MIN_VAL, PLAYER_HEAL_MAX_VAL)
	healthDiff := PLAYER_HEALTH - currPlayerHealth

	if healthDiff >= healVal {
		currPlayerHealth += healVal
	} else {
		currPlayerHealth = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttackVal := MONSTER_ATTACK_MIN_DMG
	maxAttackVal := MONSTER_ATTACK_MAX_DMG

	dmgVal := generateRandBetween(minAttackVal, maxAttackVal)
	currMonsterHealth -= dmgVal
}

func GetHealthAmounts()(int,int){
	return currPlayerHealth, currMonsterHealth
}
func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
