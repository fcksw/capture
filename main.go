package main

import (
	"capture/global"
	"capture/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	//1、加载配置文件，使用viper
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
	
	//2、加载 ORM, 使用gorm
	if err := InitDb(); err != nil {
		log.Fatal(err)
	}

	//3、加载gin，其中包括路由
	if err := InitGin(); err != nil {
		log.Fatal(err)
	}

}



func InitConfig () error {
	conf := &global.Config{}

	vip := viper.New()
	vip.AddConfigPath("./conf")
	vip.SetConfigType("toml")
	// 环境判断
	env := os.Getenv("capture_env")
	if env == "" || env == "dev" {
		// 开发环境
		vip.SetConfigName("config-dev")
	}
	if env == "prod" {
		// 生产环境
		vip.SetConfigName("config-prod")
	}

	err := vip.ReadInConfig()
	if  err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}

	if err := vip.Unmarshal(conf); err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}
	log.Printf("[config.Init] 初始化配置成功,config=%v", conf)
	global.CaptureConfig = conf
	return nil

}


func InitDb() error {

	dbName := global.CaptureConfig.Db.DbName

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		global.CaptureConfig.Db.User,
		global.CaptureConfig.Db.Password,
		global.CaptureConfig.Db.Host,
		global.CaptureConfig.Db.Port,
		dbName,)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 全局禁止表名复数
			SingularTable: true,
		},
		// 日志等级
		Logger: logger.Default.LogMode(logger.Info),
	})	

	if err != nil {
		log.Printf("[InitGorm] err = %v", err)
		return err
	}
	global.DbMapper[dbName] = db
	return nil
}


func InitGin() error {
	engin := gin.Default()
	htmlInit(engin)
	// 初始化路由
	routerInit(engin)

	//启动gin
	err := engin.Run(fmt.Sprintf(":%v", global.CaptureConfig.Gin.Port))
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



func htmlInit(server *gin.Engine) {
	// 静态资源
	server.StaticFS("/static", http.Dir("./static"))
	server.StaticFS("/views", http.Dir("./views"))
	// HTML模板加载
	server.LoadHTMLGlob("views/*")
	// 404页面
	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
}