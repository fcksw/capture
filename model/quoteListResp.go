package model

import (
	"encoding/json"
	"fmt"
)

type ItemInfo struct {
	Items []StockQuoteDailyInfo  `json:"list"`
	ItemsSize	int32			`json:"count"`
}

type QuoteListResp struct {
	Data	ItemInfo		`json:"data"`
	ErrorCode int8			`json:"error_code"`
	ErrorDescription string	`json:"error_description"`
}



func (quete *QuoteListResp) GetQuoteData(b []byte) (items []StockQuoteDailyInfo){
	resp := &QuoteListResp{}
	fmt.Println(string(b))
	json.Unmarshal(b, resp)
	items = resp.Data.Items
	return
}