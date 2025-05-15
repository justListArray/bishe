package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"` //string 换成结构体
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func callKimiAPI(apiKey string, request ChatRequest) string {
	// API 地址
	apiURL := "https://api.moonshot.cn/v1/chat/completions"

	// 将请求数据序列化为 JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// 打印响应内容
	fmt.Println(string(body))
	return string(body)
}

func UseKimi1(c *gin.Context) {
	c.HTML(http.StatusOK, "ai/index.html", nil)
}

func UseKimi(c *gin.Context) {
	var request struct {
		Content string `json:"content"`
	}
	apiKey := "sk-yCNUCISVnuKvdJcve8voHG5RU13urKGJR9IDQuF65s5l9h8u"
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	content := request.Content
	modelKimi := ChatRequest{
		Model: "moonshot-v1-8k",
		Messages: []Message{
			{Role: "user", Content: content},
		}, //更换
	}
	response := callKimiAPI(apiKey, modelKimi)
	c.JSON(http.StatusOK, response)
}
