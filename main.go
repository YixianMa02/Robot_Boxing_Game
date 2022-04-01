package main

import (
	"encoding/json"
	"fightApp/common"
	"fightApp/logic"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func startGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./resource/play.html")
}

func showRule(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./resource/rule.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("rps"))
	calcAndOutput(logic.PlayRound(playerChoice), w)
}

func playFight(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("lhd"))
	calcAndOutput(logic.PlayFightRound(playerChoice), w)
}

func resetGame() {
	common.PLAYERLIFE = common.TOTALLIFE
	common.NPCLIFE = common.TOTALLIFE
}

func calcAndOutput(result common.Round, w http.ResponseWriter) {
	if result.Winner == 1 {
		common.NPCLIFE -= 1
	}
	if result.Winner == 2 {
		common.PLAYERLIFE -= 1
	}
	if result.Winner == 5 {
		if result.ComputerChoice == "Computer chose HEAVY ATTACK" {
			common.NPCLIFE -= 2
			common.PLAYERLIFE -= 2
		} else if result.ComputerChoice == "Computer chose LIGHT ATTACK" {
			common.NPCLIFE -= 1
			common.PLAYERLIFE -= 1
		}
	}
	log.Println(result.ComputerChoice)
	result.PlayerLife = common.PLAYERLIFE
	result.NpcLife = common.NPCLIFE
	log.Println("NPC LIFE: ", common.NPCLIFE, "---PLAYER LIFE: ", common.PLAYERLIFE)
	echoJson(result, w)

	if common.NPCLIFE <= 0 || common.PLAYERLIFE <= 0 {
		resetGame()
		log.Println("Game Reset ----------------------------------------------")
		return
	}
}

func echoJson(result common.Round, response http.ResponseWriter) {
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(out)
}

func main() {
	//http.FileServer() function to create a handler which responds to all HTTP requests
	//with the contents of a given file system.
	fs := http.FileServer(http.Dir("./resource"))
	http.Handle("/", fs)

	http.HandleFunc("/play", playRound)
	http.HandleFunc("/fight", playFight)
	http.HandleFunc("/game", startGame)
	http.HandleFunc("/rule", showRule)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
