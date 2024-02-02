package handler

import (
	"capture/service"
	"log"
	"github.com/gin-gonic/gin"
	colly "github.com/gocolly/colly/v2"
)


func QuoteList(ctx *gin.Context) {
	service.GetStockServiceIns().RequestQuoteAndCreate(ctx)
}


func Login(ctx *gin.Context) {
	c := colly.NewCollector()

	err := c.Post("https://xueqiu.com/snb/provider/login", map[string]string{"username": "15522694328", "password": "Yang11120312!", "remember_me":"true"})
	if err != nil {
		log.Println(err)
	}
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.Visit("https://danjuanfunds.com/")
}
