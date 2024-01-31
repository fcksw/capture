package handler

import (
	"capture/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListUser(ctx *gin.Context) {
	result, err := service.GetUserServiceIns().ListUser(ctx)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, result)
}