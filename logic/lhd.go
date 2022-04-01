package logic

import (
	"fightApp/common"
	"math/rand"
	"time"
)

const (
	LA                = 0 // beats HA. (HA + 1) % 3 = 0
	HA                = 1 // beats DEF. (DEF + 1) % 3 = 2
	DEF               = 2 // beats LA. (LA + 1) % 3 = 1
	USERFIGHTWINS     = 1
	COMPUTERFIGHTWINS = 2
	FIGHTDRAW         = 3
	ATTACKBLOCKED     = 4
	DOUBLEDAMAGED     = 5
)

func PlayFightRound(playerValue int) common.Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case LA:
		computerChoice = "Computer chose LIGHT ATTACK"
	case DEF:
		computerChoice = "Computer chose DEFEND"
	case HA:
		computerChoice = "Computer chose HEAVY ATTACK"
	default:
	}

	if (playerValue == 0 && computerValue == 1) || (playerValue == 1 && computerValue == 2) {
		roundResult = "Player Lands a Hit!"
		winner = USERFIGHTWINS
	} else if (playerValue == 2 && computerValue == 0) || (playerValue == 0 && computerValue == 2) {
		roundResult = "Attack has Being Blocked!"
		winner = ATTACKBLOCKED
	} else if playerValue == 2 && computerValue == 2 {
		roundResult = "Both Player Defend! Nothing Happen..."
		winner = FIGHTDRAW
	} else if playerValue == computerValue {
		roundResult = "Double Puched! Both Player Loses Health!"
		winner = DOUBLEDAMAGED
	} else {
		roundResult = "Computer Lands a Hit!"
		winner = COMPUTERFIGHTWINS
	}

	var result common.Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
