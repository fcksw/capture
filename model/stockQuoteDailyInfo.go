package model

import (
	"time"
)

type StockQuoteDailyInfo struct {
	Symbol			string			`json:"symbol"`
	Name 			string			`json:"name"`
	CurrentPrice	float32			`json:"current"`
	Chg				float32			`json:"chg"`
	Percent			float32			`json:"percent"`
	CurrentYearPercent	float32		`json:"current_year_percent"`
	Volume			float32			`json:"volume"`
	Amount			float32			`json:"amount"`
	TurnoverRate	float32			`json:"turnover_rate"`
	PeTtm			float32			`json:"pe_ttm"`
	DividendYield	float32			`json:"dividend_yield"`
	MarketCapital	float32			`json:"market_capital"`
	CreatedAt		time.Time		`json:"created_at"`
	UpdatedAt		time.Time		`json:"updated_at"`
}