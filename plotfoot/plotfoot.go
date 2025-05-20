package plotfoot

import (
	"fmt"
	"footballsys/models"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotFoot(player models.Player) (file string) {
	// 示例球员数据
	// player := models.Player{
	// 	Id:         1,
	// 	Name:       "John Doe",
	// 	Goal:       10,
	// 	Pass:       50,
	// 	Tackle:     20,
	// 	Foul:       5,
	// 	Yellowcard: 2,
	// 	Redcard:    1,
	// }

	// 创建一个新的图表
	p := plot.New()

	// 设置图表标题
	p.Title.Text = fmt.Sprintf("Player Analysis: %s", player.Name)

	// 创建一个条形图数据
	labels := []string{"Goals", "Passes", "Tackles", "Fouls", "Yellow Cards", "Red Cards"}
	values := []float64{
		float64(player.Goal),
		float64(player.Pass),
		float64(player.Tackle),
		float64(player.Foul),
		float64(player.Yellowcard),
		float64(player.Redcard),
	}

	// 创建条形图
	barPlot, err := plotter.NewBarChart(plotter.Values(values), vg.Points(40))
	if err != nil {
		log.Fatalf("Failed to create bar chart: %v", err)
	}
	barPlot.LineStyle.Color = plotutil.Color(0)
	barPlot.LineStyle.Width = vg.Length(2)

	// 添加条形图到图表
	p.Add(barPlot)

	// 添加X轴标签
	p.NominalX(labels...)

	// 保存图表到文件
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "static/image/"+player.Name+"_analysis.png"); err != nil {
		log.Fatalf("Failed to save plot: %v", err)
	}

	fmt.Println("Player analysis chart saved to player_analysis.png")
	imgPath := "static/image/" + player.Name + "_analysis.png"
	if err := p.Save(10*vg.Inch, 6*vg.Inch, imgPath); err != nil {
		log.Fatalf("Failed to save plot: %v", err)
	}

	fmt.Println("Player analysis chart saved to", imgPath)
	return imgPath // 返回图片的相对路径
}
