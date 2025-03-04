package main

import (
	"fmt"
	"net/http"
	"io"
)

func GetFirstHalfStats2() {

	url := "https://free-api-live-football-data.p.rapidapi.com/football-get-match-firstHalf-stats?eventid=4621624"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "bcc44939b6msh2d153ca702afc90p1073cbjsnb6b701e7b3a8")
	req.Header.Add("x-rapidapi-host", "free-api-live-football-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}