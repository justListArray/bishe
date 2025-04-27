package models

import (
	"time"
)

// PredictionResponse 是整个JSON响应的顶级结构体
type PredictionResponse struct {
	Predictions Predictions `gorm:"embedded"`
	League      League      `gorm:"embedded"`
	Teams       Teams       `gorm:"embedded"`
	Comparison  Comparison  `gorm:"embedded"`
	H2H         []H2H       `gorm:"foreignKey:FixtureID"`
}

// Predictions 包含比赛预测信息
type Predictions struct {
	Winner struct {
		ID      uint   `gorm:"column:id"`
		Name    string `gorm:"column:name"`
		Comment string `gorm:"column:comment"`
	} `gorm:"embedded"`
	WinOrDraw bool   `gorm:"column:win_or_draw"`
	UnderOver string `gorm:"column:under_over"`
	Goals     struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	Advice  string `gorm:"column:advice"`
	Percent struct {
		Home string `gorm:"column:home"`
		Draw string `gorm:"column:draw"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
}

// League 包含联赛信息
type League struct {
	ID      uint   `gorm:"primaryKey;column:id"`
	Name    string `gorm:"column:name"`
	Country string `gorm:"column:country"`
	Logo    string `gorm:"column:logo"`
	Flag    string `gorm:"column:flag"`
	Season  int    `gorm:"column:season"`
}

// Teams 包含主队和客队信息
type Teams struct {
	Home Team `gorm:"embedded"`
	Away Team `gorm:"embedded"`
}

// Team 包含球队基本信息和统计数据
type Team struct {
	ID    uint   `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Logo  string `gorm:"column:logo"`
	Last5 struct {
		Played uint   `gorm:"column:played"`
		Form   string `gorm:"column:form"`
		Att    string `gorm:"column:att"`
		Def    string `gorm:"column:def"`
		Goals  struct {
			For struct {
				Total   uint    `gorm:"column:total"`
				Average float64 `gorm:"column:average"`
			} `gorm:"embedded"`
			Against struct {
				Total   uint    `gorm:"column:total"`
				Average float64 `gorm:"column:average"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
	League struct {
		Form     string `gorm:"column:form"`
		Fixtures struct {
			Played struct {
				Home  uint `gorm:"column:home"`
				Away  uint `gorm:"column:away"`
				Total uint `gorm:"column:total"`
			} `gorm:"embedded"`
			Wins struct {
				Home  uint `gorm:"column:home"`
				Away  uint `gorm:"column:away"`
				Total uint `gorm:"column:total"`
			} `gorm:"embedded"`
			Draws struct {
				Home  uint `gorm:"column:home"`
				Away  uint `gorm:"column:away"`
				Total uint `gorm:"column:total"`
			} `gorm:"embedded"`
			Loses struct {
				Home  uint `gorm:"column:home"`
				Away  uint `gorm:"column:away"`
				Total uint `gorm:"column:total"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
		Goals struct {
			For struct {
				Total struct {
					Home  uint `gorm:"column:home"`
					Away  uint `gorm:"column:away"`
					Total uint `gorm:"column:total"`
				} `gorm:"embedded"`
				Average struct {
					Home  string `gorm:"column:home"`
					Away  string `gorm:"column:away"`
					Total string `gorm:"column:total"`
				} `gorm:"embedded"`
				Minute struct {
					_0_15 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_16_30 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_31_45 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_46_60 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_61_75 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_76_90 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_91_105 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_106_120 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
				} `gorm:"embedded"`
				UnderOver struct {
					_0_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_1_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_2_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_3_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_4_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
				} `gorm:"embedded"`
			} `gorm:"embedded"`
			Against struct {
				Total struct {
					Home  uint `gorm:"column:home"`
					Away  uint `gorm:"column:away"`
					Total uint `gorm:"column:total"`
				} `gorm:"embedded"`
				Average struct {
					Home  string `gorm:"column:home"`
					Away  string `gorm:"column:away"`
					Total string `gorm:"column:total"`
				} `gorm:"embedded"`
				Minute struct {
					_0_15 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_16_30 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_31_45 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_46_60 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_61_75 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_76_90 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_91_105 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
					_106_120 struct {
						Total      uint   `gorm:"column:total"`
						Percentage string `gorm:"column:percentage"`
					} `gorm:"embedded"`
				} `gorm:"embedded"`
				UnderOver struct {
					_0_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_1_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_2_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_3_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
					_4_5 struct {
						Over  uint `gorm:"column:over"`
						Under uint `gorm:"column:under"`
					} `gorm:"embedded"`
				} `gorm:"embedded"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
	Biggest struct {
		Streak struct {
			Wins  uint `gorm:"column:wins"`
			Draws uint `gorm:"column:draws"`
			Loses uint `gorm:"column:loses"`
		} `gorm:"embedded"`
		Wins struct {
			Home string `gorm:"column:home"`
			Away string `gorm:"column:away"`
		} `gorm:"embedded"`
		Loses struct {
			Home string `gorm:"column:home"`
			Away string `gorm:"column:away"`
		} `gorm:"embedded"`
		Goals struct {
			For struct {
				Home uint `gorm:"column:home"`
				Away uint `gorm:"column:away"`
			} `gorm:"embedded"`
			Against struct {
				Home uint `gorm:"column:home"`
				Away uint `gorm:"column:away"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
	CleanSheet struct {
		Home  uint `gorm:"column:home"`
		Away  uint `gorm:"column:away"`
		Total uint `gorm:"column:total"`
	} `gorm:"embedded"`
	FailedToScore struct {
		Home  uint `gorm:"column:home"`
		Away  uint `gorm:"column:away"`
		Total uint `gorm:"column:total"`
	} `gorm:"embedded"`
	Penalty struct {
		Scored struct {
			Total      uint   `gorm:"column:total"`
			Percentage string `gorm:"column:percentage"`
		} `gorm:"embedded"`
		Missed struct {
			Total      uint   `gorm:"column:total"`
			Percentage string `gorm:"column:percentage"`
		} `gorm:"embedded"`
		Total uint `gorm:"column:total"`
	} `gorm:"embedded"`
	Cards struct {
		Yellow struct {
			_0_15 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_16_30 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_31_45 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_46_60 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_61_75 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_76_90 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_91_105 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_106_120 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
		Red struct {
			_0_15 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_16_30 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_31_45 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_46_60 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_61_75 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_76_90 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_91_105 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
			_106_120 struct {
				Total      uint   `gorm:"column:total"`
				Percentage string `gorm:"column:percentage"`
			} `gorm:"embedded"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
}

// Comparison 包含主队和客队的对比数据
type Comparison struct {
	Form struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	Att struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	Def struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	PoissonDistribution struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	H2H struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	Goals struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
	Total struct {
		Home string `gorm:"column:home"`
		Away string `gorm:"column:away"`
	} `gorm:"embedded"`
}

// H2H 包含历史交锋记录
type H2H struct {
	FixtureID uint `gorm:"primaryKey;column:fixture_id"`
	Fixture   struct {
		ID        uint      `gorm:"column:id"`
		Referee   string    `gorm:"column:referee"`
		Timezone  string    `gorm:"column:timezone"`
		Date      time.Time `gorm:"column:date"`
		Timestamp int64     `gorm:"column:timestamp"`
		Periods   struct {
			First  int64 `gorm:"column:first"`
			Second int64 `gorm:"column:second"`
		} `gorm:"embedded"`
		Venue struct {
			ID   uint   `gorm:"column:id"`
			Name string `gorm:"column:name"`
			City string `gorm:"column:city"`
		} `gorm:"embedded"`
		Status struct {
			Long    string `gorm:"column:long"`
			Short   string `gorm:"column:short"`
			Elapsed int    `gorm:"column:elapsed"`
			Extra   string `gorm:"column:extra"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
	League struct {
		ID        uint   `gorm:"column:id"`
		Name      string `gorm:"column:name"`
		Country   string `gorm:"column:country"`
		Logo      string `gorm:"column:logo"`
		Flag      string `gorm:"column:flag"`
		Season    int    `gorm:"column:season"`
		Round     string `gorm:"column:round"`
		Standings bool   `gorm:"column:standings"`
	} `gorm:"embedded"`
	Teams struct {
		Home struct {
			ID     uint   `gorm:"column:id"`
			Name   string `gorm:"column:name"`
			Logo   string `gorm:"column:logo"`
			Winner bool   `gorm:"column:winner"`
		} `gorm:"embedded"`
		Away struct {
			ID     uint   `gorm:"column:id"`
			Name   string `gorm:"column:name"`
			Logo   string `gorm:"column:logo"`
			Winner bool   `gorm:"column:winner"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
	Goals struct {
		Home uint `gorm:"column:home"`
		Away uint `gorm:"column:away"`
	} `gorm:"embedded"`
	Score struct {
		Halftime struct {
			Home uint `gorm:"column:home"`
			Away uint `gorm:"column:away"`
		} `gorm:"embedded"`
		Fulltime struct {
			Home uint `gorm:"column:home"`
			Away uint `gorm:"column:away"`
		} `gorm:"embedded"`
		Extratime struct {
			Home uint `gorm:"column:home"`
			Away uint `gorm:"column:away"`
		} `gorm:"embedded"`
		Penalty struct {
			Home uint `gorm:"column:home"`
			Away uint `gorm:"column:away"`
		} `gorm:"embedded"`
	} `gorm:"embedded"`
}
