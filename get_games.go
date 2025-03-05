package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Response struct {
	Response []League `json:"response"`
}

type League struct {
	Name    string   `json:"name"`
	Matches []Match `json:"matches"`
}

type Match struct {
	Home Team `json:"home"`
	Away Team `json:"away"`
}

type Team struct {
	Name string `json:"name"`
}

func getGames(date string) {

	url := "https://free-api-live-football-data.p.rapidapi.com/football-get-matches-by-date-and-league?date=" + date

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "bcc44939b6msh2d153ca702afc90p1073cbjsnb6b701e7b3a8")
	req.Header.Add("x-rapidapi-host", "free-api-live-football-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	var data Response
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, league := range data.Response {
		fmt.Println("League:", league.Name)
		for _, match := range league.Matches {
			fmt.Printf("Match: %s vs %s\n", match.Home.Name, match.Away.Name)
		}
	}

}