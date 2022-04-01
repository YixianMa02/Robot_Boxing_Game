package logic

import (
	"fightApp/common"
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

func PlayRound(playerValue int) common.Round {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	if playerValue == computerValue {
		roundResult = "It's a Draw"
		winner = DRAW
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player Wins! Player Lands a Hit!"
		winner = PLAYERWINS
	} else {
		roundResult = "Computer Wins! Computer Lands a Hit!"
		winner = COMPUTERWINS
	}

	var result common.Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result
}
