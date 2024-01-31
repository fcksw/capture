package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)


type Config struct {
	Gin		*Gin  `json:"gin"`
	Mysql	*Mysql   `json:"mysql"`
	Redis	*Redis	 `json:"redis"`
}


type Gin struct {
	Port	int64 	`json:"port"`
}

type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	DbName   string `json:"dbname"`
}


type Redis struct {
	Port	int64	`json:"port"`
	Host	string	`json:"host"`
}

//全局配置
var CaptureConfig *Config


//根据db不同区分出不同的Mapper
var DbMapper = make(map[string]*gorm.DB)


func InitResource() {
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
		panic(err)
	}
  	//2.获取所有值
  	fmt.Println("all settings: ", vip.AllSettings())
	
	if err := vip.Unmarshal(&CaptureConfig); err != nil {
		log.Printf("[config.Init] err = %v", err)
		panic(err)
	}
	log.Printf("[config.Init] 初始化配置成功,config=%v", CaptureConfig)
	log.Printf("[config.Init] 初始化配置成功,config=%v", &CaptureConfig)


}

