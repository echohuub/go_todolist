package routers

import (
	"github.com/gin-gonic/gin"
	"go_todo/controller"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "static")

	engine.GET("/", controller.IndexHandler)

	v1Group := engine.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodoHandler)

		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoListController)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodoController)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodoController)
	}
	return engine
}
