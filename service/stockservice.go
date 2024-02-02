package service

import (
	"capture/dao"
	"capture/model"
	"context"
	"sync"
)

var StockServiceIns *StockService
var StockSrvOnce sync.Once

type StockService struct {

}

func GetStockServiceIns() *StockService{
	StockSrvOnce.Do(func() {
		StockServiceIns = &StockService{}
	})
	return StockServiceIns
}


func (stockService *StockService)RequestQuoteAndCreate(ctx context.Context) {
	stockDao := dao.NewStockDao(ctx)
	b, _ := stockDao.RequestXueqiuQuoteList();
	quote := &model.QuoteListResp{}
	items := quote.GetQuoteData(b)
	for _, item := range items{
		stockDao.InsertStock(&item)
	}
}