package model

import "time"


type UserInfo struct {
	// gorm.Model
	ID				uint		`gorm:"column:id;PRIMARY_KEY"`
	Uid				uint		`gorm:"column:uid"`
	Telephone		string		`gorm:"column:telephone"`
	TradeAccountNo	string		`gorm:"column:trade_account_no"`
	TradePassword	string		`gorm:"column:trade_password"`
	Idcard			string		`gorm:"column:idcard"`
	Realname		string		`gorm:"column:realname"`
	Password		string		`gorm:"column:password"`
	Snow			string		`gorm:"column:snow"`
	TradeSnow		string		`gorm:"column:trade_snow"`
	CreateIp		string		`gorm:"column:created_ip"`
	Custno			string		`gorm:"column:custno"`
	Flavor			string		`gorm:"column:flavor"`
	Gender			int8		`gorm:"column:gender"`
	Status 			string		`gorm:"column:status"`
	ClientId		string		`gorm:"column:client_id"`
	CerType			string		`gorm:"column:cer_type"`
	RegDate			string		`gorm:"column:reg_date"`
	OpenDate		string		`gorm:"column:open_date"`
	CreateTime		time.Time		`gorm:"column:create_time"`
	UpdateTime		time.Time		`gorm:"column:update_time"`
}