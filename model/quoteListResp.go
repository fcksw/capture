package model

import (
	"capture/util"
	"encoding/json"
	"fmt"
	"time"
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
	//TODO 为什么使用  for _, item range items {}   item.TradeDate = util.StandardFmt(time.Now()) 无法赋值成功
	// 循环过程中的item 是元素的拷贝，所以这种方式修改的是拷贝的值，而不能修改元素的值
	for i := range items {		
		items[i].TradeDate = util.StandardFmt(time.Now())
	}
	return
}