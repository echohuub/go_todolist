package main

import (
	"go_todo/dao"
	"go_todo/models"
	"go_todo/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	dao.DB.AutoMigrate(&models.Todo{})

	engine := routers.SetupRouter()
	engine.Run()
}
