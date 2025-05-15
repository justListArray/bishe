package match

import (
	"fmt"
	"footballsys/controllers/base"
	"footballsys/models"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PerdictController struct {
	base.BaseController
}

func (per PerdictController) Perdict(fixture int, c *gin.Context) gin.H {
	// fixture = 198772 //strReq队伍对应的id（fixture） ,我们通过team_id 查询team的比赛，如果想要预测，就调用Predict
	url := "https://v3.football.api-sports.io/predictions?fixture=" + strconv.Itoa(fixture)

	response := per.BaseController.PreOperation(c, url)
	var strPerdiction models.PResponse
	responseData := response["response"].([]interface{})[0].(map[string]interface{})
	strPerdiction.Predictions.Winner.ID = int(responseData["predictions"].(map[string]interface{})["winner"].(map[string]interface{})["id"].(float64))
	strPerdiction.Predictions.Winner.Name = responseData["predictions"].(map[string]interface{})["winner"].(map[string]interface{})["name"].(string)
	strPerdiction.Predictions.Winner.Comment = responseData["predictions"].(map[string]interface{})["winner"].(map[string]interface{})["comment"].(string)

	strPerdiction.League.ID = int(responseData["league"].(map[string]interface{})["id"].(float64))
	strPerdiction.League.Name = responseData["league"].(map[string]interface{})["name"].(string)
	strPerdiction.League.Logo = responseData["league"].(map[string]interface{})["logo"].(string)

	strPerdiction.Teams.Home.ID = int(responseData["teams"].(map[string]interface{})["home"].(map[string]interface{})["id"].(float64))
	strPerdiction.Teams.Home.Name = responseData["teams"].(map[string]interface{})["home"].(map[string]interface{})["name"].(string)
	strPerdiction.Teams.Home.Logo = responseData["teams"].(map[string]interface{})["home"].(map[string]interface{})["logo"].(string)

	strPerdiction.Teams.Away.ID = int(responseData["teams"].(map[string]interface{})["away"].(map[string]interface{})["id"].(float64))
	strPerdiction.Teams.Away.Name = responseData["teams"].(map[string]interface{})["away"].(map[string]interface{})["name"].(string)
	strPerdiction.Teams.Away.Logo = responseData["teams"].(map[string]interface{})["away"].(map[string]interface{})["logo"].(string)

	keyInfo := gin.H{
		"Winner": strPerdiction.Predictions.Winner,
		"League": strPerdiction.League,
		"Teams": gin.H{
			"Home": strPerdiction.Teams.Home,
			"Away": strPerdiction.Teams.Away,
		},
	}
	return keyInfo
}

func (per PerdictController) QueryTeamID(name string, c *gin.Context) (fixture int) { //通过teamname找到teamid
	encodedTeamName := url.QueryEscape(name)
	url1 := "https://v3.football.api-sports.io/teams?name=" + encodedTeamName
	//log.Println(name)
	log.Printf("url:%s", url1)
	response := per.BaseController.PreOperation(c, url1)
	log.Println(response)

	responseData := response["response"].([]interface{})[0].(map[string]interface{})
	team := responseData["team"].(map[string]interface{})
	teamId := team["id"].(float64)
	// c.JSON(http.StatusOK, gin.H{
	// 	"MESSAGE": teamId,
	// })
	log.Printf("teamid%f", teamId)
	return int(teamId) //id
}

// 通过teamid找到下一次的fixtureid
func (per PerdictController) QueryTeamNextFixture(teamid int, c *gin.Context) (fixture int) {
	teamID := teamid
	season := "2023"
	baseURL := "https://v3.football.api-sports.io/fixtures" //换成下场比赛的api
	params := url.Values{}
	params.Add("team", strconv.Itoa(teamID))
	params.Add("season", season)
	url1 := baseURL + "?" + params.Encode()
	response := per.BaseController.PreOperation(c, url1)
	// fmt.Println("Response Body:", string(body))

	// err := json.Unmarshal(body, &response)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if errors, ok := response["errors"].(map[string]interface{}); ok {
	// 	if seasonError, ok := errors["season"].(string); ok {
	// 		fmt.Println("Error:", seasonError)
	// 		return
	// 	}
	// }

	responseData := response["response"].([]interface{})

	// if len(response) == 0 {
	// 	fmt.Println("No fixtures found for the specified team and season.")
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"respose": responseData,
	// })

	// 打印接下来几场比赛的信息
	for _, fixtureData := range responseData {
		fixtureArray := fixtureData.(map[string]interface{})
		fixtureID := int(fixtureArray["fixture"].(map[string]interface{})["id"].(float64))
		homeTeam := fixtureArray["teams"].(map[string]interface{})["home"].(map[string]interface{})["name"].(string)
		awayTeam := fixtureArray["teams"].(map[string]interface{})["away"].(map[string]interface{})["name"].(string)

		fmt.Printf("Fixture ID: %d\n", fixtureID)
		fmt.Printf("Home Team: %s\n", homeTeam)
		fmt.Printf("Away Team: %s\n", awayTeam)
		fmt.Println("--------------------")
		fixture = fixtureID
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": responseData,
	// })
	log.Printf("fixture id:%d", fixture)
	return fixture
}

func (per PerdictController) Prediction1(c *gin.Context) {
	c.HTML(http.StatusOK, "match/predict.html", nil)
}
func (per PerdictController) Prediction(c *gin.Context) {
	var response struct {
		Teamname string `json:"name"`
	}
	//完成我们通过team_id 查询team的比赛，如果想要预测，就调用Predict
	//下一场：
	if err := c.ShouldBindJSON(&response); err != nil {
		log.Printf("Error binding query data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("队伍的名字:%v", response.Teamname)

	CopyFixture := per.QueryTeamID(response.Teamname, c)
	fixture := per.QueryTeamNextFixture(CopyFixture, c)
	fmt.Println(fixture)
	fmt.Println(CopyFixture)
	keyInfo := PerdictController.Perdict(per, fixture, c)
	c.JSON(http.StatusOK, gin.H{
		"message": keyInfo,
	})
}
