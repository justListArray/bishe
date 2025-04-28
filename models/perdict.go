package models

type PResponse struct {
	Predictions struct {
		Winner struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Comment string `json:"comment"`
		} `json:"winner"`
	} `json:"predictions"`
	League struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"league"`
	Teams struct {
		Home struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"home"`
		Away struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		} `json:"away"`
	} `json:"teams"`
}

type Fixture struct {
	Fixture struct {
		ID int `json:"id"`
	} `json:"fixture"`
	Teams struct {
		Home struct {
			Name string `json:"name"`
		} `json:"home"`
		Away struct {
			Name string `json:"name"`
		} `json:"away"`
	} `json:"teams"`
}
