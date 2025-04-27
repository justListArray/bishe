package models

type finance struct {
	Incomes  income
	Expenses expense
}

type income struct {
	Tickets       int
	SellPlayer    int
	IncreaseBouns int
	TvTransform   int
	Sponsorship   int
	Else          int
}
type expense struct {
	Tax       int
	Salary    int
	BuyPlayer int
	Else      int
}

func (finance) TableName() string {
	return "finances"
}
