package match

import (
	"bytes"
	"encoding/json"
	"fmt"
	"footballsys/models"
	"io"
	"log"
	"net/http"
)

func CallKimiAPI(apiKey string, request models.ChatRequest) {
	// API 地址
	apiURL := "https://api.moonshot.cn/v1/chat/completions"
	//序列化为json
	requestJson, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestJson)) //发送请求
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json") //请求头

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close() //?
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Println(string(body))
}

func SimulateMatch() {
	apiKey := "sk-yCNUCISVnuKvdJcve8voHG5RU13urKGJR9IDQuF65s5l9h8u" // 替换为你的 Kimi API Key
	request := models.ChatRequest{
		Model: "moonshot-v1-8k",
		Messages: []models.Message{
			{Role: "user", Content: "你好,Kimi!"},
		}, //更换
	}
	CallKimiAPI(apiKey, request)
}

// package main

// import (
//   "fmt"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "https://v3.football.api-sports.io/leagues"
//   method := "GET"

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, nil)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("x-rapidapi-key", "XxXxXxXxXxXxXxXxXxXxXxXx")
//   req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }
