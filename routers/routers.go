package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	//静态文件路径
	r.Static("/static","static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	//请求分组 v1,CRUD
	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id",controller.DeleteATodo)
	}
	return r
}
