package base

import (
	"encoding/json"
	"fmt"
	"footballsys/models"
	"io"
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
	bodyByte, err := io.ReadAll(res.Body)
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
	if errors, ok := response["errors"].(map[string]interface{}); ok {
		if seasonError, ok := errors["season"].(string); ok {
			fmt.Println("Error:", seasonError)
			return
		}
	}
	return response
}
