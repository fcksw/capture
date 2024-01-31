package health

import (
	"capture/initialize"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HealthHanler(ctx *gin.Context) {

	ctx.String(http.StatusOK, strconv.FormatInt(initialize.CaptureConfig.Gin.Port, 10))

}