package controller

import (
	"github.com/gin-gonic/gin"
	"go_todo/models"
	"net/http"
)

func IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.AddTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todo,
	})
}

func GetTodoListController(c *gin.Context) {
	todoList, err := models.FindTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoList)
}

func UpdateTodoController(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := models.FindTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodoController(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := models.FindTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	err = models.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &todo)
}
