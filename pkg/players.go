package pkg

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
)

type Player struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Number   int    `json:"number"`
}

type Response struct {
	Players []Player `json:"players"`
}

func prepareResponse() []Player {
	var players []Player

	players = append(players, Player{Name: "Lionel Messi", Position: "Forward", Number: 10})
	players = append(players, Player{Name: "Antoine Griezmann", Position: "Forward", Number: 7})
	players = append(players, Player{Name: "Sergio Busquets", Position: "Midfielder", Number: 5})
	players = append(players, Player{Name: "Frenkie de Jong", Position: "Midfielder", Number: 21})
	players = append(players, Player{Name: "Ronald Araujo", Position: "Defender", Number: 4})
	players = append(players, Player{Name: "Jordi Alba", Position: "Defender", Number: 18})

	return players
}

func Players(w http.ResponseWriter, r *http.Request) {
	var response Response
	players := prepareResponse()

	response.Players = players

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func PlayerByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerName := vars["name"]

	allPlayers := prepareResponse()

	var foundPlayer Player
	for i := 0; i < len(allPlayers); i++ {
		if strings.EqualFold(allPlayers[i].Name, playerName) {
			foundPlayer = allPlayers[i]
			break
		}
	}

	if foundPlayer.Name == "" {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(foundPlayer)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
