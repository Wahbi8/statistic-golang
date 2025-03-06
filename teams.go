package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the structure to match the JSON data
type Result struct {
	Status   string `json:"status"`
	Result struct {
		List []TeamX `json:"list"`
	} `json:"response"`
}

type TeamX struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func GetAllTeams() {
	url := "https://free-api-live-football-data.p.rapidapi.com/football-get-list-all-team?leagueid=42"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("x-rapidapi-key", "bcc44939b6msh2d153ca702afc90p1073cbjsnb6b701e7b3a8")
	req.Header.Add("x-rapidapi-host", "free-api-live-football-data.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var data Result
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print team names and IDs
	for _, team := range data.Result.List {
		fmt.Printf("%s: %d\n", team.Name, team.ID)
	}
}
