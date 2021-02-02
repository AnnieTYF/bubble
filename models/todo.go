package models

import (
	"bubble/dao"
)

//待办事项结构体
type Todo struct{
	ID int `json:"name"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
//model层包含关于todo模型的所有增删改查
func CreateATodo(todo *Todo)(err error){
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList()(todoList []*Todo, err error)  {
	if err = dao.DB.Find(&todoList).Error;err != nil{
		return nil,err
	}
	return
}

func GetATodo(id string)(todo *Todo, err error)  {
	todo = new(Todo)
	err = dao.DB.Where("id=?",id).First(&todo).Error
	if err != nil{
		return nil,err
	}
	return
}

func UpdateATodo(todo *Todo)(err error)  {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id string)(err error){
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}