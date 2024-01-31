package initialize

import (
	"capture/model"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	_db *gorm.DB
)


/**
* gorm通过 dbresolver 作为多个database的解决方案
*/
func InitMysql() {
	mConfig := CaptureConfig.Mysql;
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mConfig.User, 
			mConfig.Password, mConfig.Host, mConfig.Port, mConfig.DbName)
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN: dsn,
	// 	DefaultStringSize: 256, // default size for string fields
  	// 	DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
  	// 	DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
  	// 	DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
  	// 	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	// }), &gorm.Config{})
	
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		//打印所有sql
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	
	_db = db
	// _ = _db.Use(dbresolver.
	// 	Register(dbresolver.Config{
	// 		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
	// 		Sources:  []gorm.Dialector{mysql.Open(pathRead)},                         // 写操作
	// 		Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)}, // 读操作
	// 		Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均衡策略
	// 	}))

	//设置表与model的连接
	_db = _db.Set("gorm:table_options", "charset=utf8mb4")
	err = migrate()
	if err != nil {
		panic(err)
	}
}


func migrate() (err error) {
	err = _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.UserInfo{})
	return
}


func NewDbClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}