package common

type Round struct {
	Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	PlayerLife     int    `json:"player_life"`
	NpcLife        int    `json:"npc_life"`
}
