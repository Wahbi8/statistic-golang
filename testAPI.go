package main

import (
	"fmt"
	"net/http"
	"io"
)

func testApi() {

	url := "https://api-football-v1.p.rapidapi.com/v3/fixtures/statistics?fixture=215662&team=463"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "bcc44939b6msh2d153ca702afc90p1073cbjsnb6b701e7b3a8")
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}