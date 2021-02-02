package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {
    //连接数据库
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	//注册路由
	r := routers.SetUpRouter()
	r.Run()
}
