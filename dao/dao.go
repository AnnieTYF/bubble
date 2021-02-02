package dao

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB  //定义全局DB变量
)

func InitMySQL()(err error){
	//连接数据库
	dsn := "root:admin@tcp(localhost:3306)/db_chat?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		return
	}
	return
}
