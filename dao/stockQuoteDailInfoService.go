package dao

import (
	"capture/initialize"
	"capture/model"
	"context"
	"io"
	"net/http"
	"time"
	"gorm.io/gorm"
)

type StockDao struct{
	*gorm.DB
}

func NewStockDao(ctx context.Context) *StockDao {
	return &StockDao{initialize.NewDbClient(ctx)}
}


//一般应用于高级的transaction等操作 
func NewStockDaoByDb(db *gorm.DB) *StockDao {
	return &StockDao{db}
}


func (dao *StockDao) InsertStock(stock *model.StockQuoteDailyInfo) (tx *gorm.DB){
	tx = dao.DB.Create(stock)
	return
}


func (dao *StockDao) RequestXueqiuQuoteList() (b []byte, err error){
	url := "https://stock.xueqiu.com/v5/stock/screener/quote/list.json?page=1&size=5&order=desc&order_by=percent&market=CN&type=sh_sz"
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Accept", "*/*")
	// 会导致返回值乱码，需要去掉； 或者，对返回值进行解包（gzip）
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("Connection","keep-alive")
	req.Header.Set("Cookie","device_id=bd014b0d719e4691668d349cb6aef9ef; s=9x1ky9cqqm; cookiesu=481694487908949; Hm_lvt_1db88642e346389874251b5a1eded6e3=1704421840; snbim_minify=true; bid=323b0820df476f4f67b1866a96ba8683_lrpzri14; whistle_nohost_env=PROD/PROD; remember=1; xq_a_token=6d110e487e5b9f8414b563e147bcf511772a787f; xqat=6d110e487e5b9f8414b563e147bcf511772a787f; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOjkyNDM4NjE3NTksImlzcyI6InVjIiwiZXhwIjoxNzA4ODI4MzA3LCJjdG0iOjE3MDY3NTYyNDY0NDUsImNpZCI6ImQ5ZDBuNEFadXAifQ.gLTz_Wt04NWaWO2qU7VnRBXZqjBqirbuaJTjev0IKtBbbJ0nh2xUlgwmtlkyP4ZQ2CCuJZqY1tdiaNn3E_gDlv854M8WHcVXn4Kiqo04H5f8l5iGDn7U3UTuSwAlc-xxeik75iTTT2K5h35WWVaNNuOZ6XNVqQbDPLbeyavDQnwKHbW1sFbLOO8RuEp4p0kyxGonn928r0dWGzl-Ghm_lbMEhCCw1_vAAjk1gu6fHwVZnu0TpoLGR1DkLhhoCMCcPBi0HZgnzQe5QDuYZVqQN6EccptCV35c0fq4FSoJmfqZBoeWWGY7oS8_ztXThmREkYIh8MeV-RN3d8ORs0PXuw; xq_r_token=19e80712bbc4747fb6a24e7bca69d046515b6305; xq_is_login=1; u=9243861759; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1706756277")
	req.Header.Set("Host", "https://xueqiu.com")
	req.Header.Set("Referer", "https://xueqiu.com/")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"")
	req.Header.Set("sec-ch-ua-platform","macOS")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req);
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	return bytes, nil
}