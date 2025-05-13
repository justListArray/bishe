package base

import (
	"encoding/json"
	"fmt"
	"footballsys/models"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct {
	db *gorm.DB
}

func (base BaseController) PreOperation(c *gin.Context, url string) (body map[string]interface{}) {
	client := &http.Client{}
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-rapidapi-key", models.ApiKey)
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("Error: API request failed with status code %d", res.StatusCode)
		bodyByte, _ := io.ReadAll(res.Body)
		log.Printf("Response body: %s", bodyByte)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API request failed"})
		return nil
	}
	bodyByte, err := io.ReadAll(res.Body)
	log.Println(string(bodyByte))
	if err != nil {
		fmt.Println(err)
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"respose": string(bodyByte),
	// })
	var response map[string]interface{}
	err = json.Unmarshal(bodyByte, &response)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("message:%+v", response)
	if errors, ok := response["errors"].(map[string]interface{}); ok {
		if seasonError, ok := errors["season"].(string); ok {
			fmt.Println("Error:", seasonError)
			return
		}
	}
	return response
}
