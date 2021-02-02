package controller
import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK,"index.html",nil)
}

func CreateATodo(context *gin.Context) {
	var todo models.Todo
	context.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK,gin.H{
			"code":200,
			"msg":"success",
			"data":todo,
		})
	}
}

func GetTodoList(context *gin.Context) {
    todoList,err := models.GetTodoList()
	if err != nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK,todoList)
	}
}

func UpdateATodo(context *gin.Context) {
	id,ok := context.Params.Get("id")
	if !ok{
		context.JSON(http.StatusOK,gin.H{"error":"无效id"})
		return
	}
	todo,err := models.GetATodo(id)
	if err != nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	context.BindJSON(&todo)
    err = models.UpdateATodo(todo)
	if err != nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(context *gin.Context) {
	id,ok := context.Params.Get("id")
	if !ok{
		context.JSON(http.StatusOK,gin.H{"error":"无效id"})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil{
		context.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK,"deleted")
	}

}
