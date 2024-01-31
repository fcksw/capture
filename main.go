package main

import (
	"capture/handler"
	"capture/health"
	"capture/initialize"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//1、加载配置文件，使用viper
	initialize.InitResource()
	//2、加载 ORM, 使用gorm
	initialize.InitMysql()
	//3、加载gin，其中包括路由
	if err := InitGin(); err != nil {
		log.Fatal(err)
	}
	//启动
}

func InitGin() error {
	engin := gin.Default()
	// htmlInit(engin)
	// 初始化路由
	routerInit(engin)

	//启动gin
	err := engin.Run(fmt.Sprintf(":%v", initialize.CaptureConfig.Gin.Port))
	if err != nil {
		log.Printf("[InitGin] err = %v", err)
	}
	log.Printf("[InitGin] success")
	return err

}

func routerInit(engin *gin.Engine) {
	// 测试
	// engin.GET("/ping", handler.Ping)
	// 权限重定向
	// engin.GET("/authority_render/:modelName", handler.RenderAuthority)
	// 首页重定向
	// engin.GET("/index", handler.Index)}

	//health相关
	healthGroup := engin.Group("/health")
	healthGroup.GET("", health.HealthHanler)

	// 账户相关
	accountGroup := engin.Group("/account")
	accountGroup.POST("/login", handler.Login)
	accountGroup.POST("/quit", handler.Quit)
	// 部门相关
	departGroup := engin.Group("/depart")
	departGroup.POST("/create", handler.DepartCreate)
	departGroup.DELETE("/del/:dep_id", handler.DepartDel)
	departGroup.POST("/edit", handler.DepartEdit)
	departGroup.GET("/query/:dep_id", handler.DepartQuery)
}

// func htmlInit(server *gin.Engine) {
// 	// 静态资源
// 	server.StaticFS("/static", http.Dir("./static"))
// 	server.StaticFS("/views", http.Dir("./views"))
// 	// HTML模板加载
// 	server.LoadHTMLGlob("views/*")
// 	// 404页面
// 	server.NoRoute(func(c *gin.Context) {
// 		c.HTML(404, "404.html", nil)
// 	})
// }
