package global

import "gorm.io/gorm"


type Config struct {
	Gin   `json:"gin"`
	Db    `json:"db"`
	Redis `json:"redis"`
}


type Gin struct {
	Port int64 `json:"port"`
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	DbName   string `json:"dbNname"`
}


type Redis struct {
	Port int64 `json:"prot"`
	Host     string `json:"host"`
}

//全局配置
var CaptureConfig *Config

//根据db不同区分出不同的Mapper
var DbMapper = make(map[string]*gorm.DB)





