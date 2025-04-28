package match

import (
	"encoding/json"
	"fmt"
	"footballsys/controllers/base"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PerdictController struct {
	base.BaseController
}

func (perdict PerdictController) Prediction(c *gin.Context) {
	// var perd string
	// err := c.ShouldBindJSON(&perd)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": perd,
	// })
	fixture := 198772 //处理？

	// Get all available predictions from one {fixture}
	//get("https://v3.football.api-sports.io/predictions?fixture=198772");  //test
	url := "https://v3.football.api-sports.io/predictions?fixture=" + strconv.Itoa(fixture)
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
	//fmt.Println("Response Body:", string(body))
	var response map[string]interface{} //  接收body
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})

}
