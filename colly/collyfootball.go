package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

// PlayerData 用于存储球员数据
type PlayerData struct {
	Rank   string
	Name   string
	Team   string
	Scores string
	League string
}

func main() {
	// 创建一个Collector
	c := colly.NewCollector(
		// 设置并发限制
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// 定义一个存储球员数据的切片
	var players []PlayerData

	// 爬取懂球帝射手榜页面
	c.OnHTML("table.table", func(e *colly.HTMLElement) {
		// 遍历表格中的每一行
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// 跳过表头
			if el.Index == 0 {
				return
			}

			// 提取球员数据
			rank := el.ChildText("td:nth-child(1)")
			name := el.ChildText("td:nth-child(2)")
			team := el.ChildText("td:nth-child(3)")
			scores := el.ChildText("td:nth-child(4)")
			league := el.ChildText("td:nth-child(5)")

			// 去除多余的空格
			name = strings.TrimSpace(name)
			team = strings.TrimSpace(team)
			scores = strings.TrimSpace(scores)
			league = strings.TrimSpace(league)

			// 存储球员数据
			player := PlayerData{
				Rank:   rank,
				Name:   name,
				Team:   team,
				Scores: scores,
				League: league,
			}
			players = append(players, player)
		})
	})

	// 在爬取完成后输出球员数据并保存到CSV文件
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("爬取完成，球员数据如下：")
		for _, player := range players {
			fmt.Printf("排名: %s, 球员: %s, 球队: %s, 进球数: %s, 联赛: %s\n",
				player.Rank, player.Name, player.Team, player.Scores, player.League)
		}

		// 保存到CSV文件
		saveToCSV(players)
	})

	// 在请求失败时打印错误信息
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// 启动爬虫
	err := c.Visit("https://www.dongqiudi.com/data/score")
	if err != nil {
		log.Fatalf("Visit failed: %v", err)
	}

	// 等待爬虫完成
	c.Wait()
}

// saveToCSV 将球员数据保存到CSV文件
func saveToCSV(players []PlayerData) {
	file, err := os.Create("players.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	header := []string{"Rank", "Name", "Team", "Scores", "League"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Failed to write header: %v", err)
	}

	// 写入球员数据
	for _, player := range players {
		record := []string{player.Rank, player.Name, player.Team, player.Scores, player.League}
		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write record: %v", err)
		}
	}

	fmt.Println("Data saved to players.csv")
}
