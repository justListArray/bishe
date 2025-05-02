package models

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://v3.football.api-sports.io/leagues"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-rapidapi-key", "aa8bc3a99aeedaedab0dfe40555e4386")
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	// c.JSON(http.StatusOK,gin.H{
	// 	"message":string(body),
	// })
}
